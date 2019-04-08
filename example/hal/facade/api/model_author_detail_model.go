package api

import (
	post "github.com/flexzuu/thesis/example/hal/post/api"
	user "github.com/flexzuu/thesis/example/hal/user/api"
	"github.com/leibowitz/halgo"
)

// AuthorDetailModel - Author with more info included
type AuthorDetailModel struct {
	halgo.Links
	AvgRating float64                   `json:"avgRating"`
	Embedded  AuthorDetailModelEmbedded `json:"_embedded"`
}
type AuthorDetailModelEmbedded struct {
	Posts  []post.PostModel `json:"posts"`
	Author user.UserModel   `json:"author"`
}
