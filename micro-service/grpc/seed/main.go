package main

import (
	"context"
	"log"
	"math/rand"
	"os"

	"github.com/flexzuu/thesis/micro-service/grpc/post/post"
	"github.com/flexzuu/thesis/micro-service/grpc/rating/rating"
	"github.com/flexzuu/thesis/micro-service/grpc/user/user"
	"google.golang.org/grpc"
)

// seed the services with some test data
func main() {

	postServiceAddress := os.Getenv("POST_SERVICE")
	if postServiceAddress == "" {
		log.Fatalln("please provide POST_SERVICE as env var")
	}
	userServiceAddress := os.Getenv("USER_SERVICE")
	if userServiceAddress == "" {
		log.Fatalln("please provide USER_SERVICE as env var")
	}
	ratingServiceAddress := os.Getenv("RATING_SERVICE")
	if ratingServiceAddress == "" {
		log.Fatalln("please provide RATING_SERVICE as env var")
	}

	postConn, err := grpc.Dial(postServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to post service: %v", err)
	}
	defer postConn.Close()
	postClient := post.NewPostServiceClient(postConn)

	userConn, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to post service: %v", err)
	}
	defer userConn.Close()
	userClient := user.NewUserServiceClient(userConn)

	ratingConn, err := grpc.Dial(ratingServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to post service: %v", err)
	}
	defer ratingConn.Close()
	ratingClient := rating.NewRatingServiceClient(ratingConn)
	ctx := context.Background()

	//createTestUser
	john, _ := userClient.Create(ctx, &user.CreateUserRequest{
		Email: "john@example.com",
		Name:  "John Doe",
	})
	//createTestUser
	jane, _ := userClient.Create(ctx, &user.CreateUserRequest{
		Email: "jane@example.com",
		Name:  "Jane Doe",
	})

	// create fake post
	var createPost = func(author *user.User, headline string) *post.Post {

		post, _ := postClient.Create(ctx, &post.CreatePostRequest{
			AuthorID: author.ID,
			Headline: headline,
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
		return post
	}

	// create fake ratings
	var rate = func(postID int64) {

		i := 1
		for i <= 5 /*num of fake reviews*/ {
			ratingClient.Create(ctx, &rating.CreateRatingRequest{
				PostID: postID,
				Rating: rand.Int31n(5),
			})
			i++
		}
	}

	worstBlogPost1 := createPost(john, "Worst Blog Post 1")
	worstBlogPost2 := createPost(john, "Worst Blog Post 2")
	worstBlogPost3 := createPost(john, "Worst Blog Post 3")
	bestBlogPost1 := createPost(jane, "Best Blog Post 1")
	bestBlogPost2 := createPost(jane, "Best Blog Post 2")
	bestBlogPost3 := createPost(jane, "Best Blog Post 3")

	rate(worstBlogPost1.ID)
	rate(worstBlogPost2.ID)
	rate(worstBlogPost3.ID)
	rate(bestBlogPost1.ID)
	rate(bestBlogPost2.ID)
	rate(bestBlogPost3.ID)

}
