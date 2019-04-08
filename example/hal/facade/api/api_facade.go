package api

import (
	"log"
	"net/http"
	"strconv"

	post "github.com/flexzuu/thesis/example/hal/post/api"
	rating "github.com/flexzuu/thesis/example/hal/rating/api"
	user "github.com/flexzuu/thesis/example/hal/user/api"

	"github.com/leibowitz/halgo"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// AuthorDetail - Author Detail
func AuthorDetail(c *gin.Context) {
	authorID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var author user.UserModel
	err = halgo.Navigator(userServiceAddress).
		Followf("find", halgo.P{
			"id": authorID,
		}).
		Unmarshal(&author)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var postsCollection post.PostListModel
	err = halgo.Navigator(userServiceAddress).
		Followf("find", halgo.P{
			"id": authorID,
		}).
		Follow("posts").
		Unmarshal(&postsCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var avgRating float64
	var length int
	posts := postsCollection.Embedded.Posts

	for _, p := range posts {
		var ratings rating.RatingListModel
		err = halgo.Navigator(ratingServiceAddress).
			Followf("ratings", halgo.P{
				"postId": p.Id,
			}).
			Unmarshal(&ratings)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for _, rating := range ratings.Embedded.Ratings {
			avgRating += float64(rating.Value)
		}
		length += ratings.Count
	}
	avgRating = avgRating / float64(length)
	c.JSON(http.StatusOK, AuthorDetailModel{
		AvgRating: avgRating,
		Links: halgo.Links{}.
			Self("/facade/authors/%d", authorID).
			Link("posts", "%s/posts?authorId=%d", postServiceAddress, authorID).
			Link("author", "%s/users/%d", userServiceAddress, authorID),
		Embedded: AuthorDetailModelEmbedded{
			Posts:  posts,
			Author: author,
		},
	})
	return
}

// ListPosts - List Posts
func ListPosts(c *gin.Context) {
	var postsCollection post.PostListModel
	err := halgo.Navigator(postServiceAddress).
		Follow("posts").
		Unmarshal(&postsCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "list failed").Error()})
		return
	}

	c.JSON(http.StatusOK, PostListModel{
		Count: postsCollection.Count,
		Embedded: PostListModelEmbedded{
			Posts: postsCollection.Embedded.Posts,
		},
		Links: halgo.Links{}.
			Self("/facade/posts").
			Link("posts", "%s/posts", postServiceAddress),
	})
	return
}

// PostDetail - Post Detail
func PostDetail(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post post.PostModel
	err = halgo.Navigator(postServiceAddress).
		Followf("find", halgo.P{
			"id": postID,
		}).
		Unmarshal(&post)
	if err != nil {
		log.Fatal(err)
	}
	var author user.UserModel
	err = halgo.Navigator(postServiceAddress).
		Followf("find", halgo.P{
			"id": postID,
		}).
		Follow("author").
		Unmarshal(&author)
	if err != nil {
		log.Fatal(err)
	}
	var ratings rating.RatingListModel
	err = halgo.Navigator(postServiceAddress).
		Followf("find", halgo.P{
			"id": postID,
		}).
		Follow("ratings").
		Unmarshal(&ratings)
	if err != nil {
		log.Fatal(err)
	}
	var avgRating float64

	for _, rating := range ratings.Embedded.Ratings {
		avgRating += float64(rating.Value)
	}
	avgRating = avgRating / float64(ratings.Count)

	c.JSON(http.StatusOK, PostDetailModel{
		Links: halgo.Links{}.
			Self("/facade/posts/%d", postID).
			Link("post", "%s/posts/%d", postServiceAddress, postID).
			Link("author", "%s/users/%d", userServiceAddress, author.Id),
		Embedded: PostDetailModelModelEmbedded{
			Author: author,
			Post:   post,
		},
		AvgRating: avgRating,
	})
	return
}
