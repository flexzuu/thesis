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
func (c *Client) RatingGet(ctx context.Context, id int) (*entity.Rating, error) {
	req := graphqlt.NewRequest(`
	query ratingGet($id: ID!) {
		rating: ratingGet(id: $id) {
		  id
		  postId
		  value
		}
	  }
	`)

	// set any variables
	req.Var("id", id)

	// run it and capture the response
	var respData struct {
		rating entity.Rating
	}
	if err := c.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData.rating, nil
}
func (c *Client) RatingListOfPost(ctx context.Context, postId int) ([]entity.Rating, error) {
	req := graphqlt.NewRequest(`
	query ratingListOfPost($postId: ID!) {
		ratings: ratingListOfPost(postId: $postId) {
		  id
		  postId
		  value
		}
	  }	  
	`)

	// set any variables
	req.Var("postId", postId)

	// run it and capture the response
	var respData struct {
		ratings []entity.Rating
	}
	if err := c.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return respData.ratings, nil
}
