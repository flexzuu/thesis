package api

import "github.com/leibowitz/halgo"

// UserModel - A User
type UserModel struct {
	halgo.Links
	Id int64 `json:"id"`

	Email string `json:"email,omitempty"`

	Name string `json:"name,omitempty"`
}
