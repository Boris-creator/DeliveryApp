package app

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"playground.com/geosuggest/internal/config"
	geosuggestpb "playground.com/proto/pkg/geosuggest"
)

func Start() error {
	app := New()
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}
	app.cfg = &cfg

	app.Bootstrap()

	lis, err := net.Listen("tcp", ":"+app.cfg.Port)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	s := grpc.NewServer()
	geosuggestpb.RegisterAddressSuggestServiceServer(s, app.server)

	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("serve: %w", err)
	}

	return nil
}
