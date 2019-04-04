package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leibowitz/halgo"
)

// CreateUser - Create user
func CreateUser(c *gin.Context) {

	var create CreateUserModel
	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usr, err := userRepo.Create(create.Email, create.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, UserModel{
		Id:    usr.ID,
		Email: usr.Email,
		Name:  usr.Name,
		Links: halgo.Links{}.
			Self("/users/%d", usr.ID).
			Link("posts", "%s/posts?authorId=%d", postServiceAddress, usr.ID),
	})
}

// DeleteUser - Delete user
func DeleteUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = userRepo.Delete(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetUserById - Get user by id
func GetUserById(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usr, err := userRepo.Get(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, UserModel{
		Id:    usr.ID,
		Email: usr.Email,
		Name:  usr.Name,
		Links: halgo.Links{}.
			Self("/users/%d", usr.ID).
			Link("posts", "%s/posts?authorId=%d", postServiceAddress, usr.ID),
	})
}
