package main

import (
	"context"
	"log"
	"math/rand"
	"os"

	post "github.com/flexzuu/thesis/example/graphql/post/repo/entity"
	"github.com/flexzuu/thesis/example/graphql/seed/postclient"
	"github.com/flexzuu/thesis/example/graphql/seed/ratingclient"
	"github.com/flexzuu/thesis/example/graphql/seed/userclient"
	user "github.com/flexzuu/thesis/example/graphql/user/repo/entity"
	"github.com/flexzuu/graphqlt"
)

// seed the services with some test data
func main() {

	postServiceEndpoint := os.Getenv("POST_SERVICE")
	if postServiceEndpoint == "" {
		log.Fatalln("please provide POST_SERVICE as env var")
	}
	userServiceEndpoint := os.Getenv("USER_SERVICE")
	if userServiceEndpoint == "" {
		log.Fatalln("please provide USER_SERVICE as env var")
	}
	ratingServiceEndpoint := os.Getenv("RATING_SERVICE")
	if ratingServiceEndpoint == "" {
		log.Fatalln("please provide RATING_SERVICE as env var")
	}

	postServiceClient := graphqlt.NewClient(postServiceEndpoint)
	userServiceClient := graphqlt.NewClient(userServiceEndpoint)
	ratingServiceClient := graphqlt.NewClient(ratingServiceEndpoint)

	postClient := postclient.Client{postServiceClient}
	userClient := userclient.Client{userServiceClient}
	ratingClient := ratingclient.Client{ratingServiceClient}

	ctx := context.Background()

	//createTestUser
	john, _ := userClient.Create(ctx, "john@example.com", "John Doe")
	//createTestUser
	jane, _ := userClient.Create(ctx, "jane@example.com", "Jane Doe")

	// create fake post
	var createPost = func(author *user.User, headline string) *post.Post {

		post, _ := postClient.Create(ctx, author.ID, headline, `Markdown is a lightweight markup language based on the formatting conventions that people naturally use in email.  As [John Gruber] writes on the [Markdown site][df1]
	
	> The overriding design goal for Markdown's
	> formatting syntax is to make it as readable
	> as possible. The idea is that a
	> Markdown-formatted document should be
	> publishable as-is, as plain text, without
	> looking like it's been marked up with tags
	> or formatting instructions.
	
	This text you see here is *actually* written in Markdown! To get a feel for Markdown's syntax, type some text into the left window and watch the results in the right.`)
		return post
	}

	// create fake ratings
	var rate = func(postID int) {

		i := 1
		for i <= 5 /*num of fake reviews*/ {
			ratingClient.Create(ctx, postID, rand.Intn(5))
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
