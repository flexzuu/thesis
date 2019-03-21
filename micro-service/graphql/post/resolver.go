package post

import (
	"context"
	"strconv"

	"github.com/pkg/errors"

	"github.com/flexzuu/benchmark/micro-service/graphql/post/repo"
	"github.com/flexzuu/benchmark/micro-service/graphql/post/userclient"
	"github.com/flexzuu/benchmark/micro-service/graphql/util"

	"github.com/flexzuu/benchmark/micro-service/graphql/post/repo/entity"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	PostRepo          repo.Post
	UserServiceClient userclient.Client
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) PostCreate(ctx context.Context, input PostCreateInput) (*entity.Post, error) {
	authorI, err := strconv.Atoi(input.AuthorID)
	if err != nil {
		return nil, err
	}
	author, err := r.UserServiceClient.UserGet(ctx, authorI)
	if err != nil || author == nil {
		return nil, errors.Wrap(err, "author not found")
	}
	post, err := r.PostRepo.Create(authorI, input.Headline, input.Content)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
func (r *mutationResolver) PostDelete(ctx context.Context, input PostDeleteInput) (*PostDeletePayload, error) {
	i, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, err
	}
	err = r.PostRepo.Delete(i)
	if err != nil {
		return nil, err
	}
	return &PostDeletePayload{
		DeletedPostID: input.ID,
	}, nil
}

type postResolver struct{ *Resolver }

func (r *postResolver) ID(ctx context.Context, obj *entity.Post) (string, error) {
	return util.UnmarshalID(obj.ID)
}
func (r *postResolver) AuthorID(ctx context.Context, obj *entity.Post) (string, error) {
	return util.UnmarshalID(obj.AuthorID)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) PostGet(ctx context.Context, id string) (*entity.Post, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	post, err := r.PostRepo.GetById(i)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
func (r *queryResolver) PostList(ctx context.Context) ([]entity.Post, error) {
	posts, err := r.PostRepo.List()
	if err != nil {
		return nil, err
	}
	return posts, nil
}
func (r *queryResolver) PostListOfAuthor(ctx context.Context, authorID string) ([]entity.Post, error) {
	authorI, err := strconv.Atoi(authorID)
	if err != nil {
		return nil, err
	}
	posts, err := r.PostRepo.ListOfAuthor(authorI)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
