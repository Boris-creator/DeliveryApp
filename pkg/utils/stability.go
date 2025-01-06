package stability

import (
	"context"
	"errors"
	"math"
	"time"

	"golang.org/x/time/rate"
)

type Effector[T, R any] func(ctx context.Context, arg T) (R, error)
type EffectorVoid[C context.Context] func(ctx C)

func Retry[T, R any](effector Effector[T, R], retries uint, delayInSeconds int) Effector[T, R] {
	if retries < 0 {
		panic("retries count must be non-negative")
	}
	return func(ctx context.Context, arg T) (R, error) {
		for r := 0; ; r++ {
			res, err := effector(ctx, arg)
			if err == nil || uint(r) == retries {
				return res, err
			}

			select {
			case <-time.After(time.Duration(math.Pow(1.5, float64(r))) * time.Second):
			case <-ctx.Done():
				return res, ctx.Err()
			}
		}
	}
}

var TooManyCallsError = errors.New("too many calls")

func Throttle[C context.Context](effector EffectorVoid[C], refill float64, max int) func(ctx C) error {
	var limiter = rate.NewLimiter(rate.Limit(refill), max)
	return func(ctx C) error {
		if !limiter.Allow() {
			return TooManyCallsError
		}
		effector(ctx)
		return nil
	}
}
