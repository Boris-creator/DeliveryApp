package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"playground.com/server/internal/config"
	approuter "playground.com/server/internal/router"
)

type App struct {
	cfg    *config.AppConfig
	db     *sqlx.DB
	server *fasthttp.Server
	rpc    *grpc.ClientConn
	router *approuter.Router
}

func New() *App {
	return &App{
		cfg:    &config.AppConfig{},
		db:     &sqlx.DB{},
		server: &fasthttp.Server{},
		rpc:    &grpc.ClientConn{},
		router: &approuter.Router{},
	}
}
