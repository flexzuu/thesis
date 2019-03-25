package userclient

import (
	"context"

	"github.com/flexzuu/benchmark/micro-service/graphql/user/repo/entity"
	"github.com/flexzuu/graphqlt"
)

type Client struct {
	*graphqlt.Client
}

// Client defines what and how to fetch
func (c *Client) Create(ctx context.Context, email, name string) (*entity.User, error) {
	req := graphqlt.NewRequest(`
	mutation create($email: String!, $name: String!) {
		user: userCreate(input: { email: $email, name: $name }) {
			id
			name
			email
		}
	}  
	`)

	// set any variables
	req.Var("email", email)
	req.Var("name", name)

	// run it and capture the response
	var respData struct {
		User entity.User
	}
	if err := c.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData.User, nil
}
