package main

import (
	"context"
	"log"
	"math/rand"
	"os"

	"github.com/flexzuu/benchmark/micro-service/grpc/post/post"
	"github.com/flexzuu/benchmark/micro-service/grpc/rating/rating"
	"github.com/flexzuu/benchmark/micro-service/grpc/user/user"
	"google.golang.org/grpc"
)

// seed the services with some test data
func main() {

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
		log.Fatalf("did not connect to post service: %v", err)
	}
	defer userConn.Close()
	userClient := user.NewUserServiceClient(userConn)

	ratingConn, err := grpc.Dial(ratingServiceAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to post service: %v", err)
	}
	defer ratingConn.Close()
	ratingClient := rating.NewRatingServiceClient(ratingConn)
	ctx := context.Background()

	//createTestUser
	john, _ := userClient.Create(ctx, &user.CreateUserRequest{
		Email: "test@example.com",
		Name:  "John Doe",
	})

	// create test posts
	worstBlogPost, _ := postClient.Create(ctx, &post.CreatePostRequest{
		AuthorID: john.ID,
		Headline: "Worst Blog Post",
		Content: `Markdown is a lightweight markup language based on the formatting conventions that people naturally use in email.  As [John Gruber] writes on the [Markdown site][df1]

> The overriding design goal for Markdown's
> formatting syntax is to make it as readable
> as possible. The idea is that a
> Markdown-formatted document should be
> publishable as-is, as plain text, without
> looking like it's been marked up with tags
> or formatting instructions.

This text you see here is *actually* written in Markdown! To get a feel for Markdown's syntax, type some text into the left window and watch the results in the right.`,
	})

	// create fake ratings

	i := 1
	for i <= 5 /*num of fake reviews*/ {
		ratingClient.Create(ctx, &rating.CreateRatingRequest{
			PostID: worstBlogPost.ID,
			Rating: rand.Int31n(5),
		})
		i++
	}

}
