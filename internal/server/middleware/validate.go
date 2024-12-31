package middleware

import (
	"encoding/json"
	"playground/internal/server/api"

	"github.com/go-playground/validator/v10"
	"github.com/valyala/fasthttp"
)

func Validate[Rules any](rh fasthttp.RequestHandler) fasthttp.RequestHandler {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return func(ctx *fasthttp.RequestCtx) {
		r, err := api.ReadRequest[Rules](ctx)
		if err != nil {
			api.ErrorResponse(ctx, err)
			return
		}
		err = validate.Struct(r)
		if validationErrors, isErr := err.(validator.ValidationErrors); isErr {
			errors := make(map[string]string, len(validationErrors))
			for _, validationError := range validationErrors {
				errors[validationError.Field()] = validationError.Error()
			}

			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			rw := json.NewEncoder(ctx)
			rw.Encode(struct {
				Errors  map[string]string `json:"errors"`
				Message string            `json:"message"`
			}{errors, err.Error()})
			return
		}
		rh(ctx)
	}
}
