package facade

import (
	"github.com/flexzuu/thesis/example/grpc/post/post"
	"github.com/flexzuu/thesis/example/grpc/rating/rating"
)

func RatingsDeRefSlice(rs []*rating.Rating) []rating.Rating {
	ras := make([]rating.Rating, len(rs))
	for i := range rs {
		ras[i] = *rs[i]
	}
	return ras
}

func PostsDeRefSlice(ps []*post.Post) []post.Post {
	pas := make([]post.Post, len(ps))
	for i := range ps {
		pas[i] = *ps[i]
	}
	return pas
}
