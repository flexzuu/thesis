package entity

import (
	pb "github.com/flexzuu/benchmark/micro-service/grpc/user/user"
)

type User struct {
	ID    int64
	Email string
	Name  string
}

func (p User) ToProto() *pb.User {
	return &pb.User{
		ID:    p.ID,
		Email: p.Email,
		Name:  p.Name,
	}
}
