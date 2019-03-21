package api

import "github.com/leibowitz/halgo"

// RatingListModel - a list of ratings
type RatingListModel struct {
	Count int `json:"count"`
	halgo.Links
	Embedded RatingListModelEmbedded `json:"_embedded"`
}
type RatingListModelEmbedded struct {
	Ratings []RatingModel `json:"ratings"`
}
