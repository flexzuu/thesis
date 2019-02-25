package stats

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
)

type statsService struct {
	Count int32
}

func (s statsService) RoundTrips(ctx context.Context, in *empty.Empty) (*RoundTripResponse, error) {
	return &RoundTripResponse{Count: s.Count}, nil
}

type CountRoundTrip = func()

func Register(s *grpc.Server) CountRoundTrip {
	srv := statsService{}
	RegisterStatsServer(s, srv)
	return func() {
		srv.Count++
	}
}
