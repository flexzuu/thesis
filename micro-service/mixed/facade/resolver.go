package facade

import (
	"context"
	"strconv"

	"github.com/flexzuu/benchmark/micro-service/graphql/util"
	"github.com/flexzuu/benchmark/micro-service/grpc/post/post"
	"github.com/flexzuu/benchmark/micro-service/grpc/rating/rating"
	"github.com/flexzuu/benchmark/micro-service/grpc/user/user"

	postEntity "github.com/flexzuu/benchmark/micro-service/grpc/post/repo/entity"
	ratingEntity "github.com/flexzuu/benchmark/micro-service/grpc/rating/repo/entity"
	userEntity "github.com/flexzuu/benchmark/micro-service/grpc/user/repo/entity"
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

func (r *postResolver) ID(ctx context.Context, obj *postEntity.Post) (string, error) {
	return util.UnmarshalID(obj.ID)
}
func (r *postResolver) Author(ctx context.Context, obj *postEntity.Post) (*userEntity.User, error) {
	u, err := r.UserClient.GetById(ctx, &user.GetUserRequest{
		ID: obj.AuthorID,
	})
	if err != nil {
		return nil, err
	}

	return &userEntity.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
func (r *postResolver) Ratings(ctx context.Context, obj *postEntity.Post) ([]ratingEntity.Rating, error) {
	ratings, err := r.RatingClient.ListOfPost(ctx, &rating.ListRatingsOfPostRequest{
		PostID: obj.ID,
	})
	if err != nil {
		return nil, err
	}
	ras := make([]ratingEntity.Rating, len(ratings.Ratings))
	for i, ra := range ratings.Ratings {
		ras[i] = ratingEntity.Rating{
			ID:     ra.ID,
			PostID: ra.PostID,
			Value:  ra.Value,
		}
	}
	return ras, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) PostGet(ctx context.Context, id string) (*postEntity.Post, error) {
	i, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return nil, err
	}
	p, err := r.PostClient.GetById(ctx, &post.GetPostRequest{
		ID: i,
	})
	if err != nil {
		return nil, err
	}

	return &postEntity.Post{
		ID:       p.ID,
		AuthorID: p.AuthorID,
		Headline: p.Headline,
		Content:  p.Content,
	}, nil
}
func (r *queryResolver) PostList(ctx context.Context) ([]postEntity.Post, error) {
	posts, err := r.PostClient.List(ctx, &post.ListPostsRequest{})
	if err != nil {
		return nil, err
	}
	ps := make([]postEntity.Post, len(posts.Posts))
	for i, p := range posts.Posts {
		ps[i] = postEntity.Post{
			ID:       p.ID,
			AuthorID: p.AuthorID,
			Headline: p.Headline,
			Content:  p.Content,
		}
	}
	return ps, nil
}
func (r *queryResolver) PostListOfAuthor(ctx context.Context, authorID string) ([]postEntity.Post, error) {
	authorI, err := strconv.ParseInt(authorID, 10, 0)
	if err != nil {
		return nil, err
	}
	posts, err := r.PostClient.ListOfAuthor(ctx, &post.ListPostsOfAuthorRequest{
		AuthorID: authorI,
	})
	if err != nil {
		return nil, err
	}
	ps := make([]postEntity.Post, len(posts.Posts))
	for i, p := range posts.Posts {
		ps[i] = postEntity.Post{
			ID:       p.ID,
			AuthorID: p.AuthorID,
			Headline: p.Headline,
			Content:  p.Content,
		}
	}
	return ps, nil
}
func (r *queryResolver) RatingGet(ctx context.Context, id string) (*ratingEntity.Rating, error) {
	i, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return nil, err
	}
	ra, err := r.RatingClient.GetById(ctx, &rating.GetRatingRequest{
		ID: i,
	})
	if err != nil {
		return nil, err
	}

	return &ratingEntity.Rating{
		ID:     ra.ID,
		PostID: ra.PostID,
		Value:  ra.Value,
	}, nil
}
func (r *queryResolver) RatingListOfPost(ctx context.Context, postID string) ([]ratingEntity.Rating, error) {
	postI, err := strconv.ParseInt(postID, 10, 0)
	if err != nil {
		return nil, err
	}
	ratings, err := r.RatingClient.ListOfPost(ctx, &rating.ListRatingsOfPostRequest{
		PostID: postI,
	})
	if err != nil {
		return nil, err
	}
	ras := make([]ratingEntity.Rating, len(ratings.Ratings))
	for i, ra := range ratings.Ratings {
		ras[i] = ratingEntity.Rating{
			ID:     ra.ID,
			PostID: ra.PostID,
			Value:  ra.Value,
		}
	}
	return ras, nil
}
func (r *queryResolver) UserGet(ctx context.Context, id string) (*userEntity.User, error) {
	i, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return nil, err
	}
	u, err := r.UserClient.GetById(ctx, &user.GetUserRequest{
		ID: i,
	})
	if err != nil {
		return nil, err
	}

	return &userEntity.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

type ratingResolver struct{ *Resolver }

func (r *ratingResolver) ID(ctx context.Context, obj *ratingEntity.Rating) (string, error) {
	return util.UnmarshalID(obj.ID)
}
func (r *ratingResolver) Post(ctx context.Context, obj *ratingEntity.Rating) (*postEntity.Post, error) {
	p, err := r.PostClient.GetById(ctx, &post.GetPostRequest{
		ID: obj.PostID,
	})
	if err != nil {
		return nil, err
	}

	return &postEntity.Post{
		ID:       p.ID,
		AuthorID: p.AuthorID,
		Headline: p.Headline,
		Content:  p.Content,
	}, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *userEntity.User) (string, error) {
	return util.UnmarshalID(obj.ID)
}
func (r *userResolver) Posts(ctx context.Context, obj *userEntity.User) ([]postEntity.Post, error) {
	posts, err := r.PostClient.ListOfAuthor(ctx, &post.ListPostsOfAuthorRequest{
		AuthorID: obj.ID,
	})
	if err != nil {
		return nil, err
	}
	ps := make([]postEntity.Post, len(posts.Posts))
	for i, p := range posts.Posts {
		ps[i] = postEntity.Post{
			ID:       p.ID,
			AuthorID: p.AuthorID,
			Headline: p.Headline,
			Content:  p.Content,
		}
	}
	return ps, nil
}
