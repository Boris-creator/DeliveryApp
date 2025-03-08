package order

import (
	"github.com/jmoiron/sqlx"
	"github.com/valyala/fasthttp"
	"playground.com/server/internal/api"
	"playground.com/server/internal/usecase/order"
	"playground.com/server/pkg/events"
)

type Handler struct {
	DB *sqlx.DB
}

// @BasePath		/api/v1
// @Summary		create order
// @Description	saving new order
// @Tags			orders
// @Accept			json
// @Produce		json
// @Param			SaveOrderRequest	body		SaveOrderRequest	true	"Order params"
// @Success		200					{object}	orderResource		"Details of the new order"
// @Router			/order [post]
func (apiHandler *Handler) Save(ctx *fasthttp.RequestCtx) {
	req, _ := api.ReadRequest[SaveOrderRequest](ctx)
	addr := req.Address.Address()

	o, err := order.SaveOrder(ctx, *apiHandler.DB, order.Order{
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
