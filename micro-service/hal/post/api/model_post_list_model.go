package api

import "github.com/leibowitz/halgo"

// PostListModel - a list of posts
type PostListModel struct {
	Count int `json:"count"`
	halgo.Links
	Embedded PostListModelEmbedded `json:"_embedded"`
}
type PostListModelEmbedded struct {
	Posts []PostModel `json:"posts"`
}
