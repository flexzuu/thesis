package userclient

import (
	"context"

	"github.com/flexzuu/thesis/micro-service/graphql/user/repo/entity"
	"github.com/flexzuu/graphqlt"
)

type Client struct {
	*graphqlt.Client
}

// Client defines what and how to fetch
func (c *Client) UserGet(ctx context.Context, id int) (*entity.User, error) {
	req := graphqlt.NewRequest(`
	query userGet($id: ID!) {
		user: userGet(id: $id) {
		  id
		  email
		  name
		}
	  }	  
	`)

	// set any variables
	req.Var("id", id)

	// run it and capture the response
	var respData struct {
		User entity.User
	}
	if err := c.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData.User, nil
}
