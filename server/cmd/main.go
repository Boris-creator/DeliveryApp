package main

import (
	"log"

	"playground.com/server/internal/app"
)

func main() {
	err := app.Start()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
