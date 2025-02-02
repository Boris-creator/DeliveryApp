package orders

import (
	"github.com/valyala/fasthttp"
	"playground/internal/server/api"
	"playground/internal/services/orders"
	"playground/pkg/events"
)

// @BasePath		/api/v1
// @Summary		create order
// @Description	saving new order
// @Tags			orders
// @Accept			json
// @Produce		json
// @Param			SaveOrderRequest	body		SaveOrderRequest	true	"Order params"
// @Success		200					{object}	orderResource		"Details of the new order"
// @Router			/order [post]
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

	_ = events.DefaultListeners.Dispatch("order:new", *o)

	api.JsonResponse(ctx, toResource(*o))
}
