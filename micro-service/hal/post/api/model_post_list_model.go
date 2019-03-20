package api

import "github.com/leibowitz/halgo"

// PostListModel - a list of posts
type PostListModel struct {
	halgo.Links
	Posts []PostModel `json:"posts"`
}
