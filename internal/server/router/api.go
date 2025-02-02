package router

import (
	"github.com/fasthttp/router"
	"playground/internal/server/api/address_suggest"
	"playground/internal/server/api/orders"
	"playground/internal/server/middleware"
)

func apiRoutes(r *router.Router) {
	api := r.Group("/api/v1")
	api.POST("/suggest-address", middleware.Throttle(address_suggest.HandleSuggest))
	api.POST("/order", middleware.Validate[orders.SaveOrderRequest](orders.SaveOrder))
}
