package api

import "github.com/leibowitz/halgo"

// RatingModel - A Rating
type RatingModel struct {
	halgo.Links
	Id    int64 `json:"id"`
	Value int32 `json:"value,omitempty"`
}
