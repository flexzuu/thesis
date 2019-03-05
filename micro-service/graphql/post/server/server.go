package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/flexzuu/benchmark/micro-service/graphql/post"
	"github.com/flexzuu/benchmark/micro-service/graphql/post/repo/inmemmory"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	postRepo := inmemmory.NewRepo()

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(post.NewExecutableSchema(post.Config{Resolvers: &post.Resolver{
		PostRepo: postRepo,
	}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
