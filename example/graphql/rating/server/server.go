package main

import (
	"log"
	"net/http"
	"os"

	"github.com/flexzuu/graphqlt"

	"github.com/99designs/gqlgen/handler"
	"github.com/flexzuu/thesis/example/graphql/rating"
	"github.com/flexzuu/thesis/example/graphql/rating/postclient"
	"github.com/flexzuu/thesis/example/graphql/rating/repo/inmemmory"
)

const defaultPort = "8082"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	ratingRepo := inmemmory.NewRepo()

	postServiceEndpoint := os.Getenv("POST_SERVICE")
	if postServiceEndpoint == "" {
		log.Fatalln("please provide POST_SERVICE as env var")
	}
	postServiceClient := graphqlt.NewClient(postServiceEndpoint)

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(rating.NewExecutableSchema(rating.Config{Resolvers: &rating.Resolver{
		RatingRepo:        ratingRepo,
		PostServiceClient: postclient.Client{postServiceClient},
	}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
