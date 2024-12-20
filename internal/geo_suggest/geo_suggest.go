package geo_suggest

import (
	"context"
	"log"
	"net"
	geo_suggest_pb "playground/proto/geo-suggest"

	dadata "github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/model"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"google.golang.org/grpc"
)

var api *suggest.Api

type server struct {
	geo_suggest_pb.UnimplementedAddressSuggestServiceServer
}

func (s *server) Suggest(_ context.Context, req *geo_suggest_pb.QueryRequest) (*geo_suggest_pb.SuggestResponse, error) {
	params := suggest.RequestParams{
		Query: req.Query,
		FromBound: &suggest.Bound{
			Value: model.BoundValue(geo_suggest_pb.Bound_name[int32(req.FromBound)]),
		},
		ToBound: &suggest.Bound{
			Value: model.BoundValue(geo_suggest_pb.Bound_name[int32(req.ToBound)]),
		},
	}

	result, err := api.Address(context.Background(), &params)
	if err != nil {
		return nil, err
	}

	suggestions := make([]*geo_suggest_pb.SuggestResponse_Result, 0, len(result))
	for _, address := range result {
		suggestions = append(suggestions, &geo_suggest_pb.SuggestResponse_Result{
			Value: address.Value,
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
