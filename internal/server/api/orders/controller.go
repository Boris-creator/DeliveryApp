package orders

import (
	"playground/internal/server/api"
	"playground/internal/services/orders"
	"playground/pkg/events"

	"github.com/valyala/fasthttp"
)

func SaveOrder(ctx *fasthttp.RequestCtx) {
	req, _ := api.ReadRequest[SaveOrderRequest](ctx)
	addr := req.Address.Address()
	o, err := orders.SaveOrder(ctx, orders.Order{
		Time:    req.Time,
		Address: addr,
		Comment: req.Comment,
	})
	if err != nil {
		api.ErrorResponse(ctx, err)
		return
	}

	l := events.DefaultListeners
	l.Dispatch("order:new", *o)

	api.JsonResponse(ctx, toResource(*o))
}
