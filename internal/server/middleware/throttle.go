package middleware

import (
	"errors"
	stability "playground/pkg/utils"

	"github.com/valyala/fasthttp"
)

func Throttle(rh fasthttp.RequestHandler) fasthttp.RequestHandler {
	throttled := stability.Throttle(stability.EffectorVoid[*fasthttp.RequestCtx](rh), 1000, 1000)
	return func(ctx *fasthttp.RequestCtx) {
		if err := throttled(ctx); errors.Is(err, stability.TooManyCallsError) {
			ctx.SetStatusCode(fasthttp.StatusTooManyRequests)
			return
		}
	}
}
