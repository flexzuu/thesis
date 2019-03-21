package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/flexzuu/benchmark/micro-service/graphql/user"
	"github.com/flexzuu/benchmark/micro-service/graphql/user/repo/inmemmory"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	userRepo := inmemmory.NewRepo()

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(user.NewExecutableSchema(user.Config{Resolvers: &user.Resolver{
		UserRepo: userRepo,
	}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
