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
	AuthorDetail(userClient, postClient, ratingClient)
}

func ListPosts(postClient post.PostServiceClient) {
	// shows post ids+headline
	ctx := context.Background()
	_ = ctx
}

func PostDetail(postClient post.PostServiceClient, userClient user.UserServiceClient, ratingClient rating.RatingServiceClient, postID int64) {
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
	avgRating /= float64(len(ratings.Ratings))

	fmt.Printf("%s by %s\nAVG-Rating: %.2f\n%s", post.Headline, author.Name, avgRating, post.Content)
}

func AuthorDetail(userClient user.UserServiceClient, postClient post.PostServiceClient, ratingClient rating.RatingServiceClient) {
	// author name and email
	// shows post ids+headline of author
	// global avg ratings
	ctx := context.Background()
	_ = ctx

}
