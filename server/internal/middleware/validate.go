package middleware

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/valyala/fasthttp"
	"playground.com/server/internal/api"
)

func Validate[Rules any](rh fasthttp.RequestHandler) fasthttp.RequestHandler {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return func(ctx *fasthttp.RequestCtx) {
		r, err := api.ReadRequest[Rules](ctx)
		if err != nil {
			api.ErrorResponse(ctx, err)

			return
		}
		// we could use validate.ValidateMap for validating form data;
		// but we prefer to convert form data to struct in api.ReadRequest
		err = validate.Struct(r)
		if validationErrors, isErr := err.(validator.ValidationErrors); isErr {
			errors := make(map[string]string, len(validationErrors))
			for _, validationError := range validationErrors {
				errors[validationError.Field()] = validationError.Error()
			}

			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			rw := json.NewEncoder(ctx)
			_ = rw.Encode(struct {
				Errors  map[string]string `json:"errors"`
				Message string            `json:"message"`
			}{errors, err.Error()})

			return
		}

		rh(ctx)
	}
}
