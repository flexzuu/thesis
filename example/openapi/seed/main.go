package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"

	post "github.com/flexzuu/thesis/example/openapi/post/openapi"
	postApi "github.com/flexzuu/thesis/example/openapi/post/openapi/client"
	rating "github.com/flexzuu/thesis/example/openapi/rating/openapi"
	ratingApi "github.com/flexzuu/thesis/example/openapi/rating/openapi/client"
	user "github.com/flexzuu/thesis/example/openapi/user/openapi"
	userApi "github.com/flexzuu/thesis/example/openapi/user/openapi/client"
)

// seed the services with some test data
func main() {

	postServiceAddress := os.Getenv("POST_SERVICE")
	if postServiceAddress == "" {
		log.Fatalln("please provide POST_SERVICE as env var")
	}
	postCfg := postApi.NewConfiguration()
	postCfg.BasePath = fmt.Sprintf("http://%s", postServiceAddress)
	postClient := postApi.NewAPIClient(postCfg)

	userServiceAddress := os.Getenv("USER_SERVICE")
	if userServiceAddress == "" {
		log.Fatalln("please provide USER_SERVICE as env var")
	}
	userCfg := userApi.NewConfiguration()
	userCfg.BasePath = fmt.Sprintf("http://%s", userServiceAddress)
	userClient := userApi.NewAPIClient(userCfg)

	ratingServiceAddress := os.Getenv("RATING_SERVICE")
	if ratingServiceAddress == "" {
		log.Fatalln("please provide RATING_SERVICE as env var")
	}
	ratingCfg := ratingApi.NewConfiguration()
	ratingCfg.BasePath = fmt.Sprintf("http://%s", ratingServiceAddress)
	ratingClient := ratingApi.NewAPIClient(ratingCfg)

	ctx := context.Background()

	//createTestUser
	john, _, _ := userClient.UserApi.CreateUser(ctx, user.CreateUserModel{
		Email: "john@example.com",
		Name:  "John Doe",
	})
	//createTestUser
	jane, _, _ := userClient.UserApi.CreateUser(ctx, user.CreateUserModel{
		Email: "jane@example.com",
		Name:  "Jane Doe",
	})

	// create fake post
	var createPost = func(author user.UserModel, headline string) post.PostModel {

		post, _, _ := postClient.PostApi.CreatePost(ctx, post.CreatePostModel{
			AuthorId: author.Id,
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
			ratingClient.RatingApi.CreateRating(ctx, rating.CreateRatingModel{
				PostId: postID,
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

	rate(worstBlogPost1.Id)
	rate(worstBlogPost2.Id)
	rate(worstBlogPost3.Id)
	rate(bestBlogPost1.Id)
	rate(bestBlogPost2.Id)
	rate(bestBlogPost3.Id)

}
