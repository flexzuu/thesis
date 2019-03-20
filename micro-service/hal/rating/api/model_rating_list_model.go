package api

import "github.com/leibowitz/halgo"

// RatingListModel - a list of ratings
type RatingListModel struct {
	halgo.Links
	Ratings []RatingModel `json:"ratings"`
}
