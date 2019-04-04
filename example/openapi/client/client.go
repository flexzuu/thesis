package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	postApi "github.com/flexzuu/thesis/example/openapi/post/openapi/client"
	ratingApi "github.com/flexzuu/thesis/example/openapi/rating/openapi/client"
	userApi "github.com/flexzuu/thesis/example/openapi/user/openapi/client"
)

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

	ListPosts(postClient.PostApi)
	PostDetail(postClient.PostApi, userClient.UserApi, ratingClient.RatingApi, 0)
	AuthorDetail(postClient.PostApi, userClient.UserApi, ratingClient.RatingApi, 0)
}

func ListPosts(postClient *postApi.PostApiService) {
	// shows post ids+headline
	ctx := context.Background()
	fmt.Println("----------ListPosts----------")
	// fetch posts
	posts, _, err := postClient.ListPosts(ctx, &postApi.ListPostsOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("#%d Posts:\n", len(posts.Posts))

	for _, post := range posts.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.Id)
	}
}

func PostDetail(postClient *postApi.PostApiService, userClient *userApi.UserApiService, ratingClient *ratingApi.RatingApiService, postID int64) {
	fmt.Println("----------PostDetail----------")
	// shows post (headline + content) + authorName and all ratings(avg)
	ctx := context.Background()
	// fetch post by id
	post, _, err := postClient.GetPostById(ctx, postID)
	if err != nil {
		log.Fatal(err)
	}
	author, _, err := userClient.GetUserById(ctx, post.AuthorId)
	if err != nil {
		log.Fatal(err)
	}
	ratings, _, err := ratingClient.ListRatings(ctx, post.Id)
	if err != nil {
		log.Fatal(err)
	}
	var avgRating float64

	for _, rating := range ratings.Ratings {
		avgRating += float64(rating.Value)
	}
	avgRating = avgRating / float64(len(ratings.Ratings))

	fmt.Printf("%s by %s\nAVG-Rating: %.2f\n%s\n", post.Headline, author.Name, avgRating, post.Content)
}

func AuthorDetail(postClient *postApi.PostApiService, userClient *userApi.UserApiService, ratingClient *ratingApi.RatingApiService, authorID int64) {
	// author name and email
	// shows post ids+headline of author
	// global avg ratings
	fmt.Println("----------AuthorDetail----------")
	ctx := context.Background()
	author, _, err := userClient.GetUserById(ctx, authorID)
	if err != nil {
		log.Fatal(err)
	}
	posts, _, err := postClient.ListPosts(ctx, &postApi.ListPostsOpts{
		AuthorId: optional.NewInt64(author.Id),
	})
	if err != nil {
		log.Fatal(err)
	}
	var avgRating float64
	var length int
	for _, post := range posts.Posts {
		ratings, _, err := ratingClient.ListRatings(ctx, post.Id)
		if err != nil {
			log.Fatal(err)
		}

		for _, rating := range ratings.Ratings {
			avgRating += float64(rating.Value)
		}
		length += len(ratings.Ratings)
	}
	avgRating = avgRating / float64(length)
	fmt.Printf("%s - %s\n", author.Name, author.Email)
	fmt.Printf("Total AVG-Rating: %.2f\n", avgRating)

	fmt.Printf("#%d Posts:\n", len(posts.Posts))

	for _, post := range posts.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.Id)
	}

}
