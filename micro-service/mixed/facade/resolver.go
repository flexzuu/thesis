package facade

import (
	"context"
	"strconv"

	"github.com/flexzuu/thesis/micro-service/graphql/util"
	"github.com/flexzuu/thesis/micro-service/grpc/post/post"
	"github.com/flexzuu/thesis/micro-service/grpc/rating/rating"
	"github.com/flexzuu/thesis/micro-service/grpc/user/user"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	PostClient   post.PostServiceClient
	UserClient   user.UserServiceClient
	RatingClient rating.RatingServiceClient
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
	return r.UserClient.GetById(ctx, &user.GetUserRequest{
		ID: obj.AuthorID,
	})
}
func (r *postResolver) Ratings(ctx context.Context, obj *post.Post) ([]rating.Rating, error) {
	resp, err := r.RatingClient.ListOfPost(ctx, &rating.ListRatingsOfPostRequest{
		PostID: obj.ID,
	})
	if err != nil {
		return nil, err
	}

	return RatingsDeRefSlice(resp.Ratings), nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) PostGet(ctx context.Context, id string) (*post.Post, error) {
	i, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return nil, err
	}
	return r.PostClient.GetById(ctx, &post.GetPostRequest{
		ID: i,
	})
}
func (r *queryResolver) PostList(ctx context.Context) ([]post.Post, error) {
	resp, err := r.PostClient.List(ctx, &post.ListPostsRequest{})
	if err != nil {
		return nil, err
	}
	return PostsDeRefSlice(resp.Posts), nil
}
func (r *queryResolver) PostListOfAuthor(ctx context.Context, authorID string) ([]post.Post, error) {
	authorI, err := strconv.ParseInt(authorID, 10, 0)
	if err != nil {
		return nil, err
	}
	resp, err := r.PostClient.ListOfAuthor(ctx, &post.ListPostsOfAuthorRequest{
		AuthorID: authorI,
	})
	return PostsDeRefSlice(resp.Posts), nil
}
func (r *queryResolver) RatingGet(ctx context.Context, id string) (*rating.Rating, error) {
	i, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return nil, err
	}
	return r.RatingClient.GetById(ctx, &rating.GetRatingRequest{
		ID: i,
	})
}
func (r *queryResolver) RatingListOfPost(ctx context.Context, postID string) ([]rating.Rating, error) {
	postI, err := strconv.ParseInt(postID, 10, 0)
	if err != nil {
		return nil, err
	}
	resp, err := r.RatingClient.ListOfPost(ctx, &rating.ListRatingsOfPostRequest{
		PostID: postI,
	})
	if err != nil {
		return nil, err
	}
	return RatingsDeRefSlice(resp.Ratings), nil
}
func (r *queryResolver) UserGet(ctx context.Context, id string) (*user.User, error) {
	i, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return nil, err
	}
	return r.UserClient.GetById(ctx, &user.GetUserRequest{
		ID: i,
	})
}

type ratingResolver struct{ *Resolver }

func (r *ratingResolver) ID(ctx context.Context, obj *rating.Rating) (string, error) {
	return util.UnmarshalID(obj.ID)
}
func (r *ratingResolver) Post(ctx context.Context, obj *rating.Rating) (*post.Post, error) {
	return r.PostClient.GetById(ctx, &post.GetPostRequest{
		ID: obj.PostID,
	})
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *user.User) (string, error) {
	return util.UnmarshalID(obj.ID)
}
func (r *userResolver) Posts(ctx context.Context, obj *user.User) ([]post.Post, error) {
	resp, err := r.PostClient.ListOfAuthor(ctx, &post.ListPostsOfAuthorRequest{
		AuthorID: obj.ID,
	})
	if err != nil {
		return nil, err
	}
	return PostsDeRefSlice(resp.Posts), nil
}
