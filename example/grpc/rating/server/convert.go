package main

import (
	pb "github.com/flexzuu/thesis/example/grpc/rating/rating"
	"github.com/flexzuu/thesis/example/grpc/rating/repo/entity"
)

func ToProto(r entity.Rating) *pb.Rating {
	return &pb.Rating{
		ID:     r.ID,
		PostID: r.PostID,
		Value:  r.Value,
	}
}
