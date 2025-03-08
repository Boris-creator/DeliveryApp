package main

import (
	"log"

	"playground.com/geosuggest/internal/app"
)

func main() {
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}

}
