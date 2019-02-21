//go:generate protoc -I ../post --go_out=plugins=grpc:../post ../post/post.proto

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/flexzuu/benchmark/micro-service/grpc/post/post"
	"github.com/flexzuu/benchmark/micro-service/grpc/post/repo"
	"github.com/flexzuu/benchmark/micro-service/grpc/post/repo/inmemmory"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement post.PostServiceServer
type server struct {
	postRepo repo.Post
}

// GetPost implements post.PostServiceServer
func (s *server) GetPost(ctx context.Context, in *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	p, err := s.postRepo.Get(in.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "get from repo failed")
	}
	return &pb.GetPostResponse{
		Post: p.ToProto(),
	}, nil
}

// CreatePost implements post.PostServiceServer
func (s *server) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	p, err := s.postRepo.Create(in.GetAuthorID(), in.GetHeadline(), in.GetContent())
	if err != nil {
		return nil, errors.Wrap(err, "create from repo failed")
	}
	return &pb.CreatePostResponse{
		Post: p.ToProto(),
	}, nil
}

// DeletePost implements post.PostServiceServer
func (s *server) DeletePost(ctx context.Context, in *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	err := s.postRepo.Delete(in.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "delete from repo failed")
	}
	return &pb.DeletePostResponse{}, nil
}

func main() {
	postRepo := inmemmory.NewRepo()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &server{
		postRepo,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
