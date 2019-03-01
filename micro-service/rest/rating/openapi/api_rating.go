/*
 * Rating Service
 *
 * a rating service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRating - Create rating
func CreateRating(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteRating - Delete rating
func DeleteRating(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetRatingById - Get rating by id
func GetRatingById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ListRatings - List ratings
func ListRatings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}