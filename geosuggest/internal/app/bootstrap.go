package app

import (
	dadata "github.com/ekomobile/dadata/v2"
	"playground.com/geosuggest/internal/server"
)

func (app *App) Bootstrap() {
	app.api = dadata.NewSuggestApi()
	app.server = &server.Server{
		Api: app.api,
	}
}
