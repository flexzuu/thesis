package postclient

import (
	"context"

	"github.com/flexzuu/thesis/example/graphql/post/repo/entity"
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
		Post entity.Post
	}
	if err := c.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData.Post, nil
}
func (c *Client) PostList(ctx context.Context) ([]entity.Post, error) {
	req := graphqlt.NewRequest(`
	query postList {
		posts: postList {
		  id
		  authorId
		  headline
		  content
		}
	  }
	`)

	// run it and capture the response
	var respData struct {
		Posts []entity.Post
	}
	if err := c.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return respData.Posts, nil
}
func (c *Client) PostListOfAuthor(ctx context.Context, authorId int) ([]entity.Post, error) {
	req := graphqlt.NewRequest(`
	query postListOfAuthor($authorId: ID!) {
		posts: postListOfAuthor(authorId:$authorId) {
		  id
		  authorId
		  headline
		  content
		}
	  }
	`)

	// set any variables
	req.Var("authorId", authorId)

	// run it and capture the response
	var respData struct {
		Posts []entity.Post
	}
	if err := c.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return respData.Posts, nil
}
