package addresssuggest

import (
	//"playground.com/server/internal/api"
	//apirouter "playground.com/server/internal/router"

	"fmt"

	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	geo_suggest_pb "playground.com/proto/pkg/geosuggest"
	"playground.com/server/internal/api"
)

type Handler struct {
	Rpc *grpc.ClientConn
}

func New(rpc *grpc.ClientConn) *Handler {
	return &Handler{
		Rpc: rpc,
	}
}

// TODO: add Swagger.
func (handler *Handler) Handle(ctx *fasthttp.RequestCtx) {
	req, _ := api.ReadRequest[suggestRequest](ctx)
	cl := geo_suggest_pb.NewAddressSuggestServiceClient(handler.Rpc)
	res, err := cl.Suggest(ctx, &geo_suggest_pb.QueryRequest{
		Query:     req.Query,
		FromBound: geo_suggest_pb.Bound(geo_suggest_pb.Bound_value[req.HighestToponym]),
		ToBound:   geo_suggest_pb.Bound(geo_suggest_pb.Bound_value[req.LowestToponym]),
	})

	if err != nil && status.Code(err) == codes.Unavailable {
		api.ErrorResponse(ctx, err)

		return
	}

	rs := make(api.ResourceCollection[resource], 0, len(res.GetSuggestions()))

	for _, address := range res.GetSuggestions() {
		rs = append(rs, resource{
			Value: address.GetValue(),
			Data: AddressData{
				GeoLat: address.GetData()["geo_lat"],
				GeoLon: address.GetData()["geo_lon"],
			},
		})
	}

	api.JsonResponse(ctx, rs)
}
