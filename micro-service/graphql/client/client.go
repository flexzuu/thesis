package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/flexzuu/thesis/micro-service/graphql/client/postclient"
	"github.com/flexzuu/thesis/micro-service/graphql/client/ratingclient"
	"github.com/flexzuu/thesis/micro-service/graphql/client/userclient"
	"github.com/flexzuu/graphqlt"
)

func main() {
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

	postClient := postclient.Client{postServiceClient}
	userClient := userclient.Client{userServiceClient}
	ratingClient := ratingclient.Client{ratingServiceClient}

	ListPosts(&postClient)
	PostDetail(&postClient, &userClient, &ratingClient, 0)
	AuthorDetail(&postClient, &userClient, &ratingClient, 0)
}

func ListPosts(postClient *postclient.Client) {
	// shows post ids+headline
	ctx := context.Background()
	fmt.Println("----------ListPosts----------")
	// fetch posts
	posts, err := postClient.PostList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("#%d Posts:\n", len(posts))

	for _, post := range posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.ID)
	}
}

func PostDetail(postClient *postclient.Client, userClient *userclient.Client, ratingClient *ratingclient.Client, postID int) {
	fmt.Println("----------PostDetail----------")
	// shows post (headline + content) + authorName and all ratings(avg)
	ctx := context.Background()
	// fetch post by id
	post, err := postClient.PostGet(ctx, postID)
	if err != nil {
		log.Fatal(err)
	}
	author, err := userClient.UserGet(ctx, post.AuthorID)
	if err != nil {
		log.Fatal(err)
	}
	ratings, err := ratingClient.RatingListOfPost(ctx, post.ID)
	if err != nil {
		log.Fatal(err)
	}
	var avgRating float64

	for _, rating := range ratings {
		avgRating += float64(rating.Value)
	}
	avgRating = avgRating / float64(len(ratings))

	fmt.Printf("%s by %s\nAVG-Rating: %.2f\n%s\n", post.Headline, author.Name, avgRating, post.Content)
}

func AuthorDetail(postClient *postclient.Client, userClient *userclient.Client, ratingClient *ratingclient.Client, authorID int) {
	// author name and email
	// shows post ids+headline of author
	// global avg ratings
	fmt.Println("----------AuthorDetail----------")
	ctx := context.Background()
	author, err := userClient.UserGet(ctx, authorID)
	if err != nil {
		log.Fatal(err)
	}
	posts, err := postClient.PostListOfAuthor(ctx, author.ID)
	if err != nil {
		log.Fatal(err)
	}
	var avgRating float64
	var length int
	for _, post := range posts {
		ratings, err := ratingClient.RatingListOfPost(ctx, post.ID)
		if err != nil {
			log.Fatal(err)
		}

		for _, rating := range ratings {
			avgRating += float64(rating.Value)
		}
		length += len(ratings)
	}
	avgRating = avgRating / float64(length)
	fmt.Printf("%s - %s\n", author.Name, author.Email)
	fmt.Printf("Total AVG-Rating: %.2f\n", avgRating)

	fmt.Printf("#%d Posts:\n", len(posts))

	for _, post := range posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.ID)
	}

}
