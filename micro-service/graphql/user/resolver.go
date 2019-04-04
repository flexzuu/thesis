package user

import (
	"context"
	"strconv"

	"github.com/flexzuu/thesis/micro-service/graphql/user/repo"
	"github.com/flexzuu/thesis/micro-service/graphql/user/repo/entity"
	"github.com/flexzuu/thesis/micro-service/graphql/util"
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
	user, err := r.UserRepo.Create(input.Email, input.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *mutationResolver) UserDelete(ctx context.Context, input UserDeleteInput) (*UserDeletePayload, error) {
	i, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, err
	}
	err = r.UserRepo.Delete(i)
	if err != nil {
		return nil, err
	}
	return &UserDeletePayload{
		DeletedUserID: input.ID,
	}, nil
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
	return util.UnmarshalID(obj.ID)
}
