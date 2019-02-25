//go:generate protoc -I ../facade --go_out=plugins=grpc:../facade ../facade/facade.proto

package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/flexzuu/benchmark/micro-service/grpc/facade/facade"
	"github.com/flexzuu/benchmark/micro-service/grpc/post/post"
	"github.com/flexzuu/benchmark/micro-service/grpc/rating/rating"
	"github.com/flexzuu/benchmark/micro-service/grpc/stats"
	"github.com/flexzuu/benchmark/micro-service/grpc/user/user"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	port = ":50060"
)

// server is used to implement facade.FacadeServiceServer
type server struct {
	postClient   post.PostServiceClient
	userClient   user.UserServiceClient
	ratingClient rating.RatingServiceClient
	stats.ServiceHelper
}

// ListPosts implements facade.FacadeServiceServer
func (s *server) ListPosts(ctx context.Context, in *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, err := s.postClient.List(ctx, &post.ListPostsRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "list failed")
	}
	s.Count++
	return &pb.ListPostsResponse{
		Posts: posts.Posts,
	}, nil
}

// PostDetail implements facade.FacadeServiceServer
func (s *server) PostDetail(ctx context.Context, in *pb.PostDetailRequest) (*pb.PostDetailResponse, error) {
	post, err := s.postClient.GetById(ctx, &post.GetPostRequest{
		ID: in.ID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list failed")
	}
	author, err := s.userClient.GetById(ctx, &user.GetUserRequest{
		ID: post.AuthorID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list failed")
	}
	ratings, err := s.ratingClient.ListOfPost(ctx, &rating.ListRatingsOfPostRequest{
		PostID: post.ID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list failed")
	}
	var avgRating float64

	for _, rating := range ratings.Ratings {
		avgRating += float64(rating.Value)
	}
	avgRating = avgRating / float64(len(ratings.Ratings))

	s.Count++
	return &pb.PostDetailResponse{
		Author:    author,
		Post:      post,
		AvgRating: avgRating,
	}, nil
}

// AuthorDetail implements facade.FacadeServiceServer
func (s *server) AuthorDetail(ctx context.Context, in *pb.AuthorDetailRequest) (*pb.AuthorDetailResponse, error) {
	author, err := s.userClient.GetById(ctx, &user.GetUserRequest{
		ID: in.ID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list failed")
	}
	posts, err := s.postClient.ListOfAuthor(ctx, &post.ListPostsOfAuthorRequest{
		AuthorID: author.ID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list failed")
	}
	var avgRating float64
	var length int
	for _, post := range posts.Posts {
		ratings, err := s.ratingClient.ListOfPost(ctx, &rating.ListRatingsOfPostRequest{
			PostID: post.ID,
		})
		if err != nil {
			return nil, errors.Wrap(err, "list failed")
		}

		for _, rating := range ratings.Ratings {
			avgRating += float64(rating.Value)
		}
		length += len(ratings.Ratings)
	}
	avgRating = avgRating / float64(length)
	s.Count++
	return &pb.AuthorDetailResponse{
		Posts:     posts.Posts,
		Author:    author,
		AvgRating: avgRating,
	}, nil
}

func main() {
	// Set up a connection to the server.
	userServiceAdress := os.Getenv("USER_SERVICE")
	if userServiceAdress == "" {
		log.Fatalln("please provide USER_SERVICE as env var")
	}
	connUser, err := grpc.Dial(userServiceAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to user service: %v", err)
	}
	defer connUser.Close()
	userClient := user.NewUserServiceClient(connUser)

	postServiceAdress := os.Getenv("POST_SERVICE")
	if postServiceAdress == "" {
		log.Fatalln("please provide POST_SERVICE as env var")
	}
	connPost, err := grpc.Dial(postServiceAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to post service: %v", err)
	}
	defer connPost.Close()
	postClient := post.NewPostServiceClient(connPost)

	ratingServiceAdress := os.Getenv("RATING_SERVICE")
	if ratingServiceAdress == "" {
		log.Fatalln("please provide RATING_SERVICE as env var")
	}
	connRating, err := grpc.Dial(ratingServiceAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to rating service: %v", err)
	}
	defer connRating.Close()
	ratingClient := rating.NewRatingServiceClient(connRating)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ServiceHelper := stats.ServiceHelper{Count: 0}
	pb.RegisterFacadeServiceServer(s, &server{
		postClient,
		userClient,
		ratingClient,
		ServiceHelper,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
