package ratingclient

import (
	"context"

	"github.com/flexzuu/benchmark/micro-service/graphql/rating/repo/entity"
	"github.com/flexzuu/graphqlt"
)

type Client struct {
	*graphqlt.Client
}

// Client defines what and how to fetch
func (c *Client) Create(ctx context.Context, postId int, rating int) (*entity.Rating, error) {
	req := graphqlt.NewRequest(`
	mutation create($postId: ID!, $rating: Int!) {
		rating: ratingCreate(input: { postId: $postId, rating: $rating }) {
			id
			postId
			value
		}
	}	
	`)

	// set any variables
	req.Var("postId", postId)
	req.Var("rating", rating)

	// run it and capture the response
	var respData struct {
		Rating entity.Rating
	}
	if err := c.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData.Rating, nil
}
