package entity

import (
	pb "github.com/flexzuu/benchmark/micro-service/grpc/post/post"
)

type Post struct {
	ID       int64
	AuthorID int64
	Headline string
	Content  string //markdown
}

func (p Post) ToProto() *pb.Post {
	return &pb.Post{
		ID:       p.ID,
		AuthorID: p.AuthorID,
		Content:  p.Content,
		Headline: p.Headline,
	}
}
