package api

import (
	"net/http"
	"strconv"

	"github.com/flexzuu/thesis/example/hal/rating/repo/entity"
	"github.com/gin-gonic/gin"
	"github.com/leibowitz/halgo"
)

// CreateRating - Create rating
func CreateRating(c *gin.Context) {
	var create CreateRatingModel
	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check postId
	checkRes, err := halgo.Navigator(postServiceAddress).
		Followf("find", halgo.P{"id": create.PostId}).
		Get()

	if err != nil || checkRes.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating, err := ratingRepo.Create(create.PostId, create.Rating)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, RatingModel{
		Id:    rating.ID,
		Value: rating.Value,
		Links: halgo.Links{}.
			Self("/ratings/%d", rating.ID).
			Link("post", "%s/posts/%d", postServiceAddress, rating.PostID),
	})
}

// DeleteRating - Delete rating
func DeleteRating(c *gin.Context) {
	ratingID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ratingRepo.Delete(ratingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetRatingById - Get rating by id
func GetRatingById(c *gin.Context) {
	ratingID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating, err := ratingRepo.GetById(ratingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, RatingModel{
		Id: rating.ID,
		Links: halgo.Links{}.
			Self("/ratings/%d", rating.ID).
			Link("post", "%s/posts/%d", postServiceAddress, rating.PostID),
		Value: rating.Value,
	})
}

// ListRatings - List ratings
func ListRatings(c *gin.Context) {
	var res []entity.Rating
	postID, err := strconv.ParseInt(c.Query("postId"), 10, 0)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check postId
	checkRes, err := halgo.Navigator(postServiceAddress).
		Followf("find", halgo.P{"id": postID}).
		Get()
	if err != nil || checkRes.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err = ratingRepo.ListOfPost(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ratings := make([]RatingModel, len(res))
	for i, r := range res {
		ratings[i] = RatingModel{
			Id: r.ID,
			Links: halgo.Links{}.
				Self("/ratings/%d", r.ID).
				Link("post", "%s/posts/%d", postServiceAddress, r.PostID),
			Value: r.Value,
		}
	}

	c.JSON(http.StatusOK, RatingListModel{
		Count: len(res),
		Links: halgo.Links{}.Self("/ratings").
			Link("find", "/ratings/{id}"),
		Embedded: RatingListModelEmbedded{
			Ratings: ratings,
		},
	})
	return
}
