package facade

import (
	"context"
	"strconv"

	"github.com/flexzuu/thesis/micro-service/graphql/facade/postclient"
	"github.com/flexzuu/thesis/micro-service/graphql/facade/ratingclient"
	"github.com/flexzuu/thesis/micro-service/graphql/facade/userclient"
	post "github.com/flexzuu/thesis/micro-service/graphql/post/repo/entity"
	rating "github.com/flexzuu/thesis/micro-service/graphql/rating/repo/entity"
	user "github.com/flexzuu/thesis/micro-service/graphql/user/repo/entity"
	"github.com/flexzuu/thesis/micro-service/graphql/util"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	UserServiceClient   userclient.Client
	PostServiceClient   postclient.Client
	RatingServiceClient ratingclient.Client
}

func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Rating() RatingResolver {
	return &ratingResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type postResolver struct{ *Resolver }

func (r *postResolver) ID(ctx context.Context, obj *post.Post) (string, error) {
	return util.UnmarshalID(obj.ID)
}
func (r *postResolver) Author(ctx context.Context, obj *post.Post) (*user.User, error) {
	user, err := r.UserServiceClient.UserGet(ctx, obj.AuthorID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *postResolver) Ratings(ctx context.Context, obj *post.Post) ([]rating.Rating, error) {
	ratings, err := r.RatingServiceClient.RatingListOfPost(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) PostGet(ctx context.Context, id string) (*post.Post, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	post, err := r.PostServiceClient.PostGet(ctx, i)
	if err != nil {
		return nil, err
	}

	return post, nil
}
func (r *queryResolver) PostList(ctx context.Context) ([]post.Post, error) {
	posts, err := r.PostServiceClient.PostList(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
func (r *queryResolver) PostListOfAuthor(ctx context.Context, authorID string) ([]post.Post, error) {
	authorI, err := strconv.Atoi(authorID)
	if err != nil {
		return nil, err
	}
	posts, err := r.PostServiceClient.PostListOfAuthor(ctx, authorI)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
func (r *queryResolver) RatingGet(ctx context.Context, id string) (*rating.Rating, error) {
	ratingI, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	rating, err := r.RatingServiceClient.RatingGet(ctx, ratingI)
	if err != nil {
		return nil, err
	}
	return rating, nil
}
func (r *queryResolver) RatingListOfPost(ctx context.Context, postID string) ([]rating.Rating, error) {
	postI, err := strconv.Atoi(postID)
	if err != nil {
		return nil, err
	}
	ratings, err := r.RatingServiceClient.RatingListOfPost(ctx, postI)
	if err != nil {
		return nil, err
	}
	return ratings, nil

}
func (r *queryResolver) UserGet(ctx context.Context, id string) (*user.User, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	user, err := r.UserServiceClient.UserGet(ctx, i)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type ratingResolver struct{ *Resolver }

func (r *ratingResolver) ID(ctx context.Context, obj *rating.Rating) (string, error) {
	return util.UnmarshalID(obj.ID)
}
func (r *ratingResolver) Post(ctx context.Context, obj *rating.Rating) (*post.Post, error) {
	post, err := r.PostServiceClient.PostGet(ctx, obj.PostID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *user.User) (string, error) {
	return util.UnmarshalID(obj.ID)
}
func (r *userResolver) Posts(ctx context.Context, obj *user.User) ([]post.Post, error) {
	posts, err := r.PostServiceClient.PostListOfAuthor(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
