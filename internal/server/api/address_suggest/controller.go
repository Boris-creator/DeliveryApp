package address_suggest

import (
	"log"
	"playground/internal/server/api"
	geo_suggest_pb "playground/proto/geo-suggest"

	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func HandleSuggest(ctx *fasthttp.RequestCtx) {
	req, _ := api.ReadRequest[suggestRequest](ctx)

	conn, err := grpc.NewClient(":8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cl := geo_suggest_pb.NewAddressSuggestServiceClient(conn)
	res, err := cl.Suggest(ctx, &geo_suggest_pb.QueryRequest{
		Query:     req.Query,
		FromBound: geo_suggest_pb.Bound(geo_suggest_pb.Bound_value[req.HighestToponym]),
		ToBound:   geo_suggest_pb.Bound(geo_suggest_pb.Bound_value[req.LowestToponym]),
	})

	if err != nil && status.Code(err) == codes.Unavailable {
		api.ErrorResponse(ctx, err)
		return
	}
	var rs = make(api.ResourceCollection[resource], 0, len(res.Suggestions))
	for _, address := range res.Suggestions {
		rs = append(rs, resource{
			Value: address.Value,
			Data: AddressData{
				GeoLat: address.Data["geo_lat"],
				GeoLon: address.Data["geo_lon"],
			},
		})
	}
	api.JsonResponse(ctx, rs)
}
