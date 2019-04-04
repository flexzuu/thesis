package api

import (
	"net/http"
	"strconv"

	"github.com/flexzuu/thesis/micro-service/hal/post/repo/entity"
	"github.com/gin-gonic/gin"
	"github.com/leibowitz/halgo"
)

// CreatePost - Create post
func CreatePost(c *gin.Context) {
	var create CreatePostModel
	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check authorId
	checkRes, err := halgo.Navigator(userServiceAddress).
		Followf("find", halgo.P{"id": create.AuthorId}).
		Get()
	if err != nil || checkRes.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := postRepo.Create(create.AuthorId, create.Headline, create.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, PostModel{
		Id: post.ID,
		Links: halgo.Links{}.
			Self("/posts/%d", post.ID).
			Link("author", "%s/users/%d", userServiceAddress, post.AuthorID).
			Link("ratings", "%s/ratings?postId=%d", ratingServiceAddress, post.ID),

		Headline: post.Headline,
		Content:  post.Content,
	})
}

// DeletePost - Delete post
func DeletePost(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = postRepo.Delete(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetPostById - Get post by id
func GetPostById(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := postRepo.GetById(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, PostModel{
		Id: post.ID,
		Links: halgo.Links{}.
			Self("/posts/%d", post.ID).
			Link("author", "%s/users/%d", userServiceAddress, post.AuthorID).
			Link("ratings", "%s/ratings?postId=%d", ratingServiceAddress, post.ID),
		Headline: post.Headline,
		Content:  post.Content,
	})
}

// ListPosts - List posts
func ListPosts(c *gin.Context) {
	var res []entity.Post

	authorID, err := strconv.ParseInt(c.Query("authorId"), 10, 0)
	if err != nil {
		res, err = postRepo.List()
	} else {
		//check authorId
		checkRes, err := halgo.Navigator(userServiceAddress).
			Followf("find", halgo.P{"id": authorID}).
			Get()
		if err != nil || checkRes.StatusCode != http.StatusOK {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err = postRepo.ListOfAuthor(authorID)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	posts := make([]PostModel, len(res))
	for i, p := range res {
		posts[i] = PostModel{
			Id: p.ID,
			Links: halgo.Links{}.
				Self("/posts/%d", p.ID).
				Link("author", "%s/users/%d", userServiceAddress, p.AuthorID).
				Link("ratings", "%s/ratings?postId=%d", ratingServiceAddress, p.ID),
			Headline: p.Headline,
			Content:  p.Content,
		}
	}

	c.JSON(http.StatusOK, PostListModel{
		Count: len(res),
		Links: halgo.Links{}.
			Self("/posts").
			Link("find", "/posts/{id}"),
		Embedded: PostListModelEmbedded{
			Posts: posts,
		},
	})
	return
}
