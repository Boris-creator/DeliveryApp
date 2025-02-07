package geo_suggest

import (
	"context"
	"log"
	"net"

	dadata "github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/model"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"google.golang.org/grpc"
	stability "playground/pkg/utils"
	geo_suggest_pb "playground/proto/geo-suggest"
)

var api *suggest.Api

type server struct {
	geo_suggest_pb.UnimplementedAddressSuggestServiceServer
}

var (
	retries        uint = 3
	delay               = 1
	suggestAddress      = stability.Retry[suggest.RequestParams, []*suggest.AddressSuggestion](
		func(ctx context.Context, params suggest.RequestParams) ([]*suggest.AddressSuggestion, error) {
			return api.Address(ctx, &params)
		}, retries, delay,
	)
)

func (s *server) Suggest(ctx context.Context, req *geo_suggest_pb.QueryRequest) (*geo_suggest_pb.SuggestResponse, error) {
	params := suggest.RequestParams{
		Query: req.GetQuery(),
		FromBound: &suggest.Bound{
			Value: model.BoundValue(geo_suggest_pb.Bound_name[int32(req.GetFromBound())]),
		},
		ToBound: &suggest.Bound{
			Value: model.BoundValue(geo_suggest_pb.Bound_name[int32(req.GetToBound())]),
		},
	}

	result, err := suggestAddress(ctx, params)
	if err != nil {
		return nil, err
	}

	suggestions := make([]*geo_suggest_pb.SuggestResponse_Result, 0, len(result))
	for _, address := range result {
		suggestions = append(suggestions, &geo_suggest_pb.SuggestResponse_Result{
			Value: address.Value,
			Data: map[string]string{
				"geo_lat": address.Data.GeoLat,
				"geo_lon": address.Data.GeoLon,
			},
		})
	}

	res := geo_suggest_pb.SuggestResponse{
		Suggestions: suggestions,
	}

	return &res, nil
}

func StartServer(cfg config) {
	api = dadata.NewSuggestApi()

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	geo_suggest_pb.RegisterAddressSuggestServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
