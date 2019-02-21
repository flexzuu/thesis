package entity

import (
	pb "github.com/flexzuu/benchmark/micro-service/grpc/user/user"
	validator "gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

type User struct {
	ID    int64
	Email string `validate:"required"`
	Name  string `validate:"required"`
}

func (u User) ToProto() *pb.User {
	return &pb.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}
}

func (u *User) Valid() error {
	return validate.Struct(u)
}
