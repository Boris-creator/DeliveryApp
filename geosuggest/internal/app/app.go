package app

import (
	"github.com/ekomobile/dadata/v2/api/suggest"
	"playground.com/geosuggest/internal/config"
	"playground.com/geosuggest/internal/server"
)

type App struct {
	api    *suggest.Api
	server *server.Server
	cfg    *config.Config
}

func New() *App {
	return &App{
		api:    &suggest.Api{},
		server: &server.Server{},
		cfg:    &config.Config{},
	}
}
