package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"playground.com/server/internal/config"
	"playground.com/server/internal/logger"
	"playground.com/server/internal/models"
	"playground.com/server/internal/usecase/order"
	"playground.com/server/internal/usecase/work"
	"playground.com/server/pkg/events"
)

func Start() error {
	app := New()

	defer app.Shutdown()

	logger.Init()

	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}
	app.cfg = &cfg

	dbConn, err := models.Connect(cfg)
	if err != nil {
		return fmt.Errorf("connecting database: %w", err)
	}
	app.db = dbConn

	rpcConn, err := app.startGrpcClient()
	if err != nil {
		return fmt.Errorf("did not connect grpc server: %w", err)
	}
	app.rpc = rpcConn

	app.Bootstrap()

	app.registerEventsListeners()

	if err := app.startServer(); err != nil {
		return fmt.Errorf("listenAndServe: %w", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}

func (app *App) startServer() error {
	logger.Info("server started on port %s", app.cfg.Port)
	return app.server.ListenAndServe(net.JoinHostPort(app.cfg.Host, app.cfg.Port))
}

func (app *App) startGrpcClient() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", app.cfg.GeoSuggestHost, app.cfg.GeoSuggestPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (app *App) registerEventsListeners() {
	l := events.DefaultListeners
	l.AddListener(context.TODO(), "order:new", func(e events.Event[any]) {
		o, _ := e.Payload.(order.Order)

		ws, err := models.FindNearestWorkshop(app.db, o.Address.GeoLat, o.Address.GeoLon)
		if err != nil {
			logger.Error(err)

			return
		}

		wm := models.NewWorkModel(app.db)

		wm.Model = models.Work{
			OrderId:    o.Id,
			WorkshopId: ws.Id,
			Status:     uint8(work.StatusPending),
			StartAt:    nil,
		}

		_, err = wm.Create()
		if err != nil {
			logger.Error(err)
		}
	})
	// l.Listen() // not necessary when running from main function
}
