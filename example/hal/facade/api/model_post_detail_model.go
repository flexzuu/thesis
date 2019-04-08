package api

import (
	post "github.com/flexzuu/thesis/example/hal/post/api"
	user "github.com/flexzuu/thesis/example/hal/user/api"
	"github.com/leibowitz/halgo"
)

// PostDetailModel - Post with more info included
type PostDetailModel struct {
	halgo.Links

	AvgRating float64                      `json:"avgRating"`
	Embedded  PostDetailModelModelEmbedded `json:"_embedded"`
}
type PostDetailModelModelEmbedded struct {
	Post   post.PostModel `json:"post"`
	Author user.UserModel `json:"author"`
}
