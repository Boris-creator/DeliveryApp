package app

import (
	"context"

	"github.com/ekomobile/dadata/v2/api/model"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"playground.com/geosuggest/pkg/stability"
	geosuggestpb "playground.com/proto/pkg/geosuggest"
)

func (app *App) Suggest(ctx context.Context, req *geosuggestpb.QueryRequest) (*geosuggestpb.SuggestResponse, error) {
	const retries = 3
	const delay = 1

	suggestAddress := stability.Retry[suggest.RequestParams, []*suggest.AddressSuggestion](
		func(ctx context.Context, params suggest.RequestParams) ([]*suggest.AddressSuggestion, error) {
			return app.api.Address(ctx, &params)
		}, retries, delay,
	)

	params := suggest.RequestParams{
		Query: req.GetQuery(),
		FromBound: &suggest.Bound{
			Value: model.BoundValue(geosuggestpb.Bound_name[int32(req.GetFromBound())]),
		},
		ToBound: &suggest.Bound{
			Value: model.BoundValue(geosuggestpb.Bound_name[int32(req.GetToBound())]),
		},
	}

	result, err := suggestAddress(ctx, params)
	if err != nil {
		return nil, err
	}

	suggestions := make([]*geosuggestpb.SuggestResponse_Result, 0, len(result))
	for _, address := range result {
		suggestions = append(suggestions, &geosuggestpb.SuggestResponse_Result{
			Value: address.Value,
			Data: map[string]string{
				"geo_lat": address.Data.GeoLat,
				"geo_lon": address.Data.GeoLon,
			},
		})
	}

	res := geosuggestpb.SuggestResponse{
		Suggestions: suggestions,
	}

	return &res, nil
}
