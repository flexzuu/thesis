package user

import (
	"context"
	"strconv"

	"github.com/flexzuu/benchmark/micro-service/graphql/user/repo"
	"github.com/flexzuu/benchmark/micro-service/graphql/user/repo/entity"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	UserRepo repo.User
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) UserCreate(ctx context.Context, input UserCreateInput) (*entity.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UserDelete(ctx context.Context, input UserDeleteInput) (*UserDeletePayload, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) UserGet(ctx context.Context, id string) (*entity.User, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	user, err := r.UserRepo.Get(i)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *entity.User) (string, error) {
	panic("not implemented")
}
