package server

import (
	"fmt"
	"log"

	"playground/internal/config"
	"playground/internal/models"
	"playground/internal/server/router"

	"github.com/valyala/fasthttp"
)

func StartServer() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg := config.Config
	c, err := models.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	r := router.MakeRouter()

	s := &fasthttp.Server{
		Handler: r.Handler,
		Name:    "Playground",
	}

	registerEventsListeners()

	if err := s.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil {
		log.Fatalf("error in ListenAndServe: %v", err)
	}
}
