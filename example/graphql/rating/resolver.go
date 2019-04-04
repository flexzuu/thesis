package rating

import (
	"context"
	"strconv"

	"github.com/pkg/errors"

	"github.com/flexzuu/thesis/example/graphql/rating/postclient"
	"github.com/flexzuu/thesis/example/graphql/rating/repo"
	"github.com/flexzuu/thesis/example/graphql/util"

	"github.com/flexzuu/thesis/example/graphql/rating/repo/entity"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	RatingRepo        repo.Rating
	PostServiceClient postclient.Client
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Rating() RatingResolver {
	return &ratingResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) RatingCreate(ctx context.Context, input RatingCreateInput) (*entity.Rating, error) {
	postI, err := strconv.Atoi(input.PostID)
	if err != nil {
		return nil, err
	}
	post, err := r.PostServiceClient.PostGet(ctx, postI)
	if err != nil || post == nil {
		return nil, errors.Wrap(err, "post not found")
	}
	rating, err := r.RatingRepo.Create(postI, input.Rating)
	if err != nil {
		return nil, err
	}
	return &rating, nil
}
func (r *mutationResolver) RatingDelete(ctx context.Context, input RatingDeleteInput) (*RatingDeletePayload, error) {
	i, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, err
	}
	err = r.RatingRepo.Delete(i)
	if err != nil {
		return nil, err
	}
	return &RatingDeletePayload{
		DeletedRatingID: input.ID,
	}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) RatingGet(ctx context.Context, id string) (*entity.Rating, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	rating, err := r.RatingRepo.GetById(i)
	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func (r *queryResolver) RatingListOfPost(ctx context.Context, postID string) ([]entity.Rating, error) {
	postI, err := strconv.Atoi(postID)
	if err != nil {
		return nil, err
	}
	ratings, err := r.RatingRepo.ListOfPost(postI)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

type ratingResolver struct{ *Resolver }

func (r *ratingResolver) ID(ctx context.Context, obj *entity.Rating) (string, error) {
	return util.UnmarshalID(obj.ID)
}
func (r *ratingResolver) PostID(ctx context.Context, obj *entity.Rating) (string, error) {
	return util.UnmarshalID(obj.PostID)
}
