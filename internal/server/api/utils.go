package api

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/valyala/fasthttp"
	"playground/pkg/ioutils"
)

func JsonResponse(ctx *fasthttp.RequestCtx, data any) {
	ctx.SetContentType("application/json")
	rw := json.NewEncoder(ctx)
	_ = rw.Encode(data)
}

func ReadRequest[data any](ctx *fasthttp.RequestCtx) (data, error) {
	var rq data
	ct := ctx.Request.Header.ContentType()

	if strings.HasPrefix(string(ct), "application/json") || strings.HasPrefix(string(ct), "text/plain") {
		err := json.Unmarshal(ctx.PostBody(), &rq)

		return rq, err
	}

	if strings.HasPrefix(string(ct), "multipart/form-data") {
		mf, _ := ctx.MultipartForm()
		err := ioutils.ReadFormToStruct(mf, &rq)

		return rq, err
	}

	return rq, errors.New("unsupported request content type")
}

func ErrorResponse(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	rw := json.NewEncoder(ctx)
	_ = rw.Encode(struct {
		Error string `json:"error"`
	}{err.Error()})
}

type ResourceCollection[R any] []R
