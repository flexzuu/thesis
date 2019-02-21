//go:generate protoc -I ../rating --go_out=plugins=grpc:../rating ../rating/rating.proto

package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/flexzuu/benchmark/micro-service/grpc/post/post"

	pb "github.com/flexzuu/benchmark/micro-service/grpc/rating/rating"
	"github.com/flexzuu/benchmark/micro-service/grpc/rating/repo"
	"github.com/flexzuu/benchmark/micro-service/grpc/rating/repo/inmemmory"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	port = ":50053"
)

// server is used to implement rating.RatingServiceServer
type server struct {
	ratingRepo repo.Rating
	postClient post.PostServiceClient
}

// GetRating implements rating.RatingServiceServer
func (s *server) GetById(ctx context.Context, in *pb.GetRatingRequest) (*pb.Rating, error) {
	r, err := s.ratingRepo.GetById(in.ID)
	if err != nil {
		return nil, errors.Wrap(err, "get failed")
	}
	return ToProto(r), nil
}

// ListOfPost implements rating.RatingServiceServer
func (s *server) ListOfPost(ctx context.Context, in *pb.ListRatingsOfPostRequest) (*pb.ListRatingsResponse, error) {
	rs, err := s.ratingRepo.ListOfPost(in.PostID)
	if err != nil {
		return nil, errors.Wrap(err, "list failed")
	}
	ratings := make([]*pb.Rating, len(rs))
	for i, r := range rs {
		ratings[i] = ToProto(r)
	}
	return &pb.ListRatingsResponse{
		Ratings: ratings,
	}, nil
}

// CreateRating implements rating.RatingServiceServer
func (s *server) Create(ctx context.Context, in *pb.CreateRatingRequest) (*pb.Rating, error) {
	//Validate PostID with post service
	_, err := s.postClient.GetById(ctx, &post.GetPostRequest{
		ID: in.PostID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "invalid post id")
	}

	r, err := s.ratingRepo.Create(in.PostID, in.Rating)
	if err != nil {
		return nil, errors.Wrap(err, "create failed")
	}
	return ToProto(r), nil
}

// DeleteRating implements rating.RatingServiceServer
func (s *server) Delete(ctx context.Context, in *pb.DeleteRatingRequest) (*empty.Empty, error) {
	err := s.ratingRepo.Delete(in.ID)
	if err != nil {
		return nil, errors.Wrap(err, "delete failed")
	}
	return &empty.Empty{}, nil
}

func main() {
	ratingRepo := inmemmory.NewRepo()

	conn, err := grpc.Dial(os.Getenv("POST_SERVICE"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to post service: %v", err)
	}
	defer conn.Close()
	postClient := post.NewPostServiceClient(conn)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRatingServiceServer(s, &server{
		ratingRepo,
		postClient,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
