package main

import (
	"log"
	"net/http"
	"os"

	"github.com/flexzuu/graphqlt"

	"github.com/99designs/gqlgen/handler"
	"github.com/flexzuu/thesis/micro-service/graphql/facade"
	"github.com/flexzuu/thesis/micro-service/graphql/facade/postclient"
	"github.com/flexzuu/thesis/micro-service/graphql/facade/ratingclient"
	"github.com/flexzuu/thesis/micro-service/graphql/facade/userclient"
)

const defaultPort = "8090"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userServiceEndpoint := os.Getenv("USER_SERVICE")
	if userServiceEndpoint == "" {
		log.Fatalln("please provide USER_SERVICE as env var")
	}
	userServiceClient := graphqlt.NewClient(userServiceEndpoint)

	postServiceEndpoint := os.Getenv("POST_SERVICE")
	if postServiceEndpoint == "" {
		log.Fatalln("please provide POST_SERVICE as env var")
	}
	postServiceClient := graphqlt.NewClient(postServiceEndpoint)

	ratingServiceEndpoint := os.Getenv("RATING_SERVICE")
	if ratingServiceEndpoint == "" {
		log.Fatalln("please provide RATING_SERVICE as env var")
	}
	ratingServiceClient := graphqlt.NewClient(ratingServiceEndpoint)

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(facade.NewExecutableSchema(facade.Config{Resolvers: &facade.Resolver{
		UserServiceClient:   userclient.Client{userServiceClient},
		PostServiceClient:   postclient.Client{postServiceClient},
		RatingServiceClient: ratingclient.Client{ratingServiceClient},
	}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
