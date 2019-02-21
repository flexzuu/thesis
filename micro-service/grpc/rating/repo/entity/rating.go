package entity

import (
	pb "github.com/flexzuu/benchmark/micro-service/grpc/rating/rating"
)

type Rating struct {
	ID     int64
	PostID int64
	Value  int32
}

func (r Rating) ToProto() *pb.Rating {
	return &pb.Rating{
		ID:     r.ID,
		PostID: r.PostID,
		Value:  r.Value,
	}
}
