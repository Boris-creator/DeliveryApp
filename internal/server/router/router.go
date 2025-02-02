package router

import (
	"github.com/fasthttp/router"
)

func MakeRouter() *router.Router {
	r := router.New()
	webRoutes(r)
	apiRoutes(r)

	return r
}
