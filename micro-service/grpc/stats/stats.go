package stats

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type ServiceHelper struct {
	Count int32
}

func (s *ServiceHelper) RoundTrips(ctx context.Context, in *empty.Empty) (*RoundTripResponse, error) {
	return &RoundTripResponse{Count: s.Count}, nil
}
