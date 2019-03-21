package main

import (
	"log"
	"net/http"
	"os"

	"github.com/flexzuu/graphqlt"

	"github.com/99designs/gqlgen/handler"
	"github.com/flexzuu/benchmark/micro-service/graphql/post"
	"github.com/flexzuu/benchmark/micro-service/graphql/post/repo/inmemmory"
	"github.com/flexzuu/benchmark/micro-service/graphql/post/userclient"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	postRepo := inmemmory.NewRepo()

	userServiceEndpoint := os.Getenv("USER_SERVICE")
	if userServiceEndpoint == "" {
		log.Fatalln("please provide USER_SERVICE as env var")
	}
	userServiceClient := graphqlt.NewClient(userServiceEndpoint)

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(post.NewExecutableSchema(post.Config{Resolvers: &post.Resolver{
		PostRepo:          postRepo,
		UserServiceClient: userclient.Client{userServiceClient},
	}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
