package api

import (
	post "github.com/flexzuu/thesis/example/hal/post/api"
	"github.com/leibowitz/halgo"
)

// PostListModel - a list of posts
type PostListModel struct {
	Count int `json:"count"`
	halgo.Links
	Embedded PostListModelEmbedded `json:"_embedded"`
}
type PostListModelEmbedded struct {
	Posts []post.PostModel `json:"posts"`
}
