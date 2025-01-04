package main

import (
	"playground/internal/logger"
	"playground/internal/server"
)

func main() {
	logger.Init()
	logger.Info("Starting server")
	server.StartServer()
}
