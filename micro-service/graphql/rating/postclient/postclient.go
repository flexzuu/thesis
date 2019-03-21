package postclient

import (
	"context"

	"github.com/flexzuu/benchmark/micro-service/graphql/post/repo/entity"
	"github.com/flexzuu/graphqlt"
)

type Client struct {
	*graphqlt.Client
}

// Client defines what and how to fetch
func (c *Client) PostGet(ctx context.Context, id int) (*entity.Post, error) {
	req := graphqlt.NewRequest(`
	query postGet($id: ID!) {
		post: postGet(id: $id) {
			id
			authorId
			headline
			content
		}
	}
	`)

	// set any variables
	req.Var("id", id)

	// run it and capture the response
	var respData struct {
		Data struct {
			Post entity.Post
		}
	}
	if err := c.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData.Data.Post, nil
}
