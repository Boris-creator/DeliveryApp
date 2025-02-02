package main

import (
	"log"

	geo_suggest "playground/internal/geo_suggest"
)

func main() {
	cfg, err := geo_suggest.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	geo_suggest.StartServer(cfg)
}
