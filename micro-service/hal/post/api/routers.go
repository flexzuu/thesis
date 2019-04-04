package api

import (
	"log"
	"net/http"
	"os"

	"github.com/leibowitz/halgo"

	"github.com/flexzuu/thesis/micro-service/hal/post/repo"
	"github.com/flexzuu/thesis/micro-service/hal/post/repo/inmemmory"
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
var postRepo repo.Post
var userServiceAddress string
var ratingServiceAddress string

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	postRepo = inmemmory.NewRepo()

	userServiceAddress = os.Getenv("USER_SERVICE")
	if userServiceAddress == "" {
		log.Fatalln("please provide USER_SERVICE as env var")
	}

	ratingServiceAddress = os.Getenv("RATING_SERVICE")
	if ratingServiceAddress == "" {
		log.Fatalln("please provide RATING_SERVICE as env var")
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
			Link("posts", "/posts{?authorId}").
			Link("find", "/posts/{id}"),
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
		"CreatePost",
		http.MethodPost,
		"/posts",
		CreatePost,
	},

	{
		"DeletePost",
		http.MethodDelete,
		"/posts/:id",
		DeletePost,
	},

	{
		"GetPostById",
		http.MethodGet,
		"/posts/:id",
		GetPostById,
	},

	{
		"ListPosts",
		http.MethodGet,
		"/posts",
		ListPosts,
	},
}
