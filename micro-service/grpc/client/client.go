package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/flexzuu/benchmark/micro-service/grpc/post/post"
	"github.com/flexzuu/benchmark/micro-service/grpc/rating/rating"
	"github.com/flexzuu/benchmark/micro-service/grpc/user/user"
	"google.golang.org/grpc"
)

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

	ListPosts(postClient)
	PostDetail(postClient, userClient, ratingClient, 0)
	AuthorDetail(userClient, postClient, ratingClient, 0)
}

func ListPosts(postClient post.PostServiceClient) {
	// shows post ids+headline
	ctx := context.Background()
	fmt.Println("----------ListPosts----------")
	// fetch posts
	posts, err := postClient.List(ctx, &post.ListPostsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("#%d Posts:\n", len(posts.Posts))

	for _, post := range posts.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.ID)
	}
}

func PostDetail(postClient post.PostServiceClient, userClient user.UserServiceClient, ratingClient rating.RatingServiceClient, postID int64) {
	fmt.Println("----------PostDetail----------")
	// shows post (headline + content) + authorName and all ratings(avg)
	ctx := context.Background()
	// fetch post by id
	post, err := postClient.GetById(ctx, &post.GetPostRequest{
		ID: postID,
	})
	if err != nil {
		log.Fatal(err)
	}
	author, err := userClient.GetById(ctx, &user.GetUserRequest{
		ID: post.AuthorID,
	})
	if err != nil {
		log.Fatal(err)
	}
	ratings, err := ratingClient.ListOfPost(ctx, &rating.ListRatingsOfPostRequest{
		PostID: post.ID,
	})
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

func AuthorDetail(userClient user.UserServiceClient, postClient post.PostServiceClient, ratingClient rating.RatingServiceClient, authorID int64) {
	// author name and email
	// shows post ids+headline of author
	// global avg ratings
	fmt.Println("----------AuthorDetail----------")
	ctx := context.Background()
	author, err := userClient.GetById(ctx, &user.GetUserRequest{
		ID: authorID,
	})
	if err != nil {
		log.Fatal(err)
	}
	posts, err := postClient.ListOfAuthor(ctx, &post.ListPostsOfAuthorRequest{
		AuthorID: author.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	var avgRating float64
	var length int
	for _, post := range posts.Posts {
		ratings, err := ratingClient.ListOfPost(ctx, &rating.ListRatingsOfPostRequest{
			PostID: post.ID,
		})
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
		fmt.Printf("\t%s (%d)\n", post.Headline, post.ID)
	}

}
