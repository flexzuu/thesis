package main

import (
	pb "github.com/flexzuu/benchmark/micro-service/grpc/post/post"
	"github.com/flexzuu/benchmark/micro-service/grpc/post/repo/entity"
)

func ToProto(p entity.Post) *pb.Post {
	return &pb.Post{
		ID:       p.ID,
		AuthorID: p.AuthorID,
		Content:  p.Content,
		Headline: p.Headline,
	}
}
