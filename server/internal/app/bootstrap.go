package app

import (
	"github.com/valyala/fasthttp"
	"playground.com/server/internal/router"
)

func (app *App) Bootstrap() {
	app.router = router.Make()

	app.router.Handlers.Order.DB = app.db
	app.router.Handlers.Geosuggest.Rpc = app.rpc

	app.server = &fasthttp.Server{
		Handler: app.router.Handler,
		Name:    "Playground",
	}
}
