package postclient

import (
	"context"

	"github.com/flexzuu/thesis/micro-service/graphql/post/repo/entity"
	"github.com/flexzuu/graphqlt"
)

type Client struct {
	*graphqlt.Client
}

// Client defines what and how to fetch
func (c *Client) Create(ctx context.Context, authorId int, headline, content string) (*entity.Post, error) {
	req := graphqlt.NewRequest(`
	mutation create($authorId: ID!, $headline: String!, $content: String!) {
		post: postCreate(
		  input: { authorId: $authorId, headline: $headline, content: $content }
		) {
		  id
		  authorId
		  headline
		  content
		}
	  }
	`)

	// set any variables
	req.Var("authorId", authorId)
	req.Var("headline", headline)
	req.Var("content", content)

	// run it and capture the response
	var respData struct {
		Post entity.Post
	}
	if err := c.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData.Post, nil
}
