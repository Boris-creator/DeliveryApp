package router

import (
	"github.com/fasthttp/router"
	"playground.com/server/internal/api/addresssuggest"
	"playground.com/server/internal/api/order"
)

type Handlers struct {
	Order      order.Handler
	Geosuggest addresssuggest.Handler
}

type Router struct {
	*router.Router
	Handlers Handlers
}

func Make() *Router {
	r := Router{
		Router: router.New(),
	}
	r.webRoutes()
	r.apiRoutes()

	return &r
}
