package api

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func JsonResponse(ctx *fasthttp.RequestCtx, data any) {
	ctx.SetContentType("application/json")
	rw := json.NewEncoder(ctx)
	rw.Encode(data)
}

func ReadRequest[data any](ctx *fasthttp.RequestCtx) (data, error) {
	var rq data
	err := json.Unmarshal(ctx.PostBody(), &rq)
	return rq, err
}

func ErrorResponse(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	rw := json.NewEncoder(ctx)
	rw.Encode(struct {
		Error string `json:"error"`
	}{err.Error()})
}

type ResourceCollection[R any] []R
