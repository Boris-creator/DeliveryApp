package server

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
	"playground/internal/config"
	"playground/internal/models"
	"playground/internal/server/router"
)

func StartServer() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = models.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

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
