package router

import (
	"playground.com/server/internal/api/order"
	"playground.com/server/internal/middleware"
)

func (r *Router) apiRoutes() {
	api := r.Group("/api/v1")
	api.POST("/suggest-address", middleware.Throttle(r.Handlers.Geosuggest.Handle))
	api.POST("/order", middleware.Validate[order.SaveOrderRequest](r.Handlers.Order.Save))
}
