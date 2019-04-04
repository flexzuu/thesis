package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/flexzuu/thesis/micro-service/grpc/post/post"
	"github.com/flexzuu/thesis/micro-service/grpc/rating/rating"
	"github.com/flexzuu/thesis/micro-service/grpc/user/user"
	"github.com/flexzuu/thesis/micro-service/mixed/facade"
	"google.golang.org/grpc"
)

const defaultPort = "8090"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	postServiceAdress := os.Getenv("POST_SERVICE")
	if postServiceAdress == "" {
		log.Fatalln("please provide POST_SERVICE as env var")
	}
	userServiceAdress := os.Getenv("USER_SERVICE")
	if postServiceAdress == "" {
		log.Fatalln("please provide USER_SERVICE as env var")
	}
	ratingServiceAdress := os.Getenv("RATING_SERVICE")
	if postServiceAdress == "" {
		log.Fatalln("please provide RATING_SERVICE as env var")
	}

	postConn, err := grpc.Dial(postServiceAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to post service: %v", err)
	}
	defer postConn.Close()
	postClient := post.NewPostServiceClient(postConn)

	userConn, err := grpc.Dial(userServiceAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to user service: %v", err)
	}
	defer userConn.Close()
	userClient := user.NewUserServiceClient(userConn)

	ratingConn, err := grpc.Dial(ratingServiceAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to rating service: %v", err)
	}
	defer ratingConn.Close()
	ratingClient := rating.NewRatingServiceClient(ratingConn)

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(facade.NewExecutableSchema(facade.Config{Resolvers: &facade.Resolver{
		UserClient:   userClient,
		PostClient:   postClient,
		RatingClient: ratingClient,
	}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
