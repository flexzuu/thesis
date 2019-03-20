package api

import (
	"log"
	"net/http"
	"os"

	"github.com/flexzuu/benchmark/micro-service/hal/rating/repo"
	"github.com/flexzuu/benchmark/micro-service/hal/rating/repo/inmemmory"
	"github.com/gin-gonic/gin"
	"github.com/leibowitz/halgo"
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
var postServiceAddress string

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	ratingRepo = inmemmory.NewRepo()

	postServiceAddress = os.Getenv("POST_SERVICE")
	if postServiceAddress == "" {
		log.Fatalln("please provide POST_SERVICE as env var")
	}

	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Root is the index handler.
func Root(c *gin.Context) {
	type root struct{ halgo.Links }

	r := root{
		Links: halgo.Links{}.
			Self("/").
			Link("ratings", "/ratings{?postId}").
			Link("find", "/ratings/{id}"),
	}
	c.JSON(http.StatusOK, r)
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Root,
	},

	{
		"CreateRating",
		http.MethodPost,
		"/ratings",
		CreateRating,
	},

	{
		"DeleteRating",
		http.MethodDelete,
		"/ratings/:id",
		DeleteRating,
	},

	{
		"GetRatingById",
		http.MethodGet,
		"/ratings/:id",
		GetRatingById,
	},

	{
		"ListRatings",
		http.MethodGet,
		"/ratings",
		ListRatings,
	},
}
