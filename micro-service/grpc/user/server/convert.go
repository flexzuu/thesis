package main

import (
	"github.com/flexzuu/benchmark/micro-service/grpc/user/repo/entity"
	pb "github.com/flexzuu/benchmark/micro-service/grpc/user/user"
)

func ToProto(u entity.User) *pb.User {
	return &pb.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}
}
