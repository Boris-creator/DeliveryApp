package middleware

import (
	"errors"

	stability "playground/pkg/utils"

	"github.com/valyala/fasthttp"
)

func Throttle(rh fasthttp.RequestHandler) fasthttp.RequestHandler {
	var refill float64 = 1000

	lim := 1000
	throttled := stability.Throttle(stability.EffectorVoid[*fasthttp.RequestCtx](rh), refill, lim)

	return func(ctx *fasthttp.RequestCtx) {
		if err := throttled(ctx); errors.Is(err, stability.TooManyCallsError) {
			ctx.SetStatusCode(fasthttp.StatusTooManyRequests)

			return
		}
	}
}
