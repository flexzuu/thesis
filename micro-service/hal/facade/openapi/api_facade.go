/*
 * Facade Service
 *
 * a facade service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"strconv"

	"github.com/antihax/optional"
	post "github.com/flexzuu/benchmark/micro-service/rest/post/openapi/client"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// AuthorDetail - Author Detail
func AuthorDetail(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	author, _, err := userServiceClient.UserApi.GetUserById(ctx, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "list failed").Error()})
		return
	}
	posts, _, err := postServiceClient.PostApi.ListPosts(ctx, &post.ListPostsOpts{
		AuthorId: optional.NewInt64(author.Id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "list failed").Error()})
		return
	}

	var avgRating float64
	var length int
	for _, post := range posts.Posts {
		ratings, _, err := ratingServiceClient.RatingApi.ListRatings(ctx, post.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "list failed").Error()})
			return
		}

		for _, rating := range ratings.Ratings {
			avgRating += float64(rating.Value)
		}
		length += len(ratings.Ratings)
	}
	avgRating = avgRating / float64(length)

	c.JSON(http.StatusOK, AuthorDetailModel{
		Author:    author,
		Posts:     posts,
		AvgRating: avgRating,
	})
	return
}

// ListPosts - List Posts
func ListPosts(c *gin.Context) {
	ctx := context.Background()

	posts, _, err := postServiceClient.PostApi.ListPosts(ctx, &post.ListPostsOpts{
		AuthorId: optional.EmptyInt64(),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "list failed").Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
	return
}

// PostDetail - Post Detail
func PostDetail(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	post, _, err := postServiceClient.PostApi.GetPostById(ctx, postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "list failed").Error()})
		return
	}
	author, _, err := userServiceClient.UserApi.GetUserById(ctx, post.AuthorId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "list failed").Error()})
		return
	}
	ratings, _, err := ratingServiceClient.RatingApi.ListRatings(ctx, post.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "list failed").Error()})
		return
	}
	var avgRating float64

	for _, rating := range ratings.Ratings {
		avgRating += float64(rating.Value)
	}
	avgRating = avgRating / float64(len(ratings.Ratings))

	c.JSON(http.StatusOK, PostDetailModel{
		Post:      post,
		Author:    author,
		AvgRating: avgRating,
	})
	return
}