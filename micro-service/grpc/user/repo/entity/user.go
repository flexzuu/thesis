package entity

import (
	validator "gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

type User struct {
	ID    int64
	Email string `validate:"required"`
	Name  string `validate:"required"`
}

func (u *User) Valid() error {
	return validate.Struct(u)
}
