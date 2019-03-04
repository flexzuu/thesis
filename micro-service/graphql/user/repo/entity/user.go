package entity

import (
	validator "gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email" validate:"required"`
	Name  string `json:"name" validate:"required"`
}

func (u *User) Valid() error {
	return validate.Struct(u)
}
