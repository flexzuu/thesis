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
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	postApi "github.com/flexzuu/thesis/example/rest/post/openapi/client"
	"github.com/flexzuu/thesis/example/rest/rating/repo"
	"github.com/flexzuu/thesis/example/rest/rating/repo/inmemmory"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

//dependencies
var ratingRepo repo.Rating
var postServiceClient *postApi.APIClient

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	ratingRepo = inmemmory.NewRepo()

	postServiceAddress := os.Getenv("POST_SERVICE")
	if postServiceAddress == "" {
		log.Fatalln("please provide POST_SERVICE as env var")
	}
	postCfg := postApi.NewConfiguration()
	postCfg.BasePath = fmt.Sprintf("http://%s", postServiceAddress)
	postServiceClient = postApi.NewAPIClient(postCfg)

	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			router.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			router.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"CreateRating",
		strings.ToUpper("Post"),
		"/rating",
		CreateRating,
	},

	{
		"DeleteRating",
		strings.ToUpper("Delete"),
		"/rating/:id",
		DeleteRating,
	},

	{
		"GetRatingById",
		strings.ToUpper("Get"),
		"/rating/:id",
		GetRatingById,
	},

	{
		"ListRatings",
		strings.ToUpper("Get"),
		"/rating",
		ListRatings,
	},
}