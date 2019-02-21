//go:generate protoc -I ../post --go_out=plugins=grpc:../post ../post/post.proto

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/flexzuu/benchmark/micro-service/grpc/post/post"
	"github.com/flexzuu/benchmark/micro-service/grpc/post/repo"
	"github.com/flexzuu/benchmark/micro-service/grpc/post/repo/inmemmory"
	"github.com/golang/protobuf/ptypes/empty"
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
func (s *server) GetById(ctx context.Context, in *pb.GetPostRequest) (*pb.Post, error) {
	p, err := s.postRepo.Get(in.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "get from repo failed")
	}
	return p.ToProto(), nil
}

// CreatePost implements post.PostServiceServer
func (s *server) Create(ctx context.Context, in *pb.CreatePostRequest) (*pb.Post, error) {
	//TODO: Validate AuthorID with user service
	p, err := s.postRepo.Create(in.GetAuthorID(), in.GetHeadline(), in.GetContent())
	if err != nil {
		return nil, errors.Wrap(err, "create from repo failed")
	}
	return p.ToProto(), nil
}

// DeletePost implements post.PostServiceServer
func (s *server) Delete(ctx context.Context, in *pb.DeletePostRequest) (*empty.Empty, error) {
	err := s.postRepo.Delete(in.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "delete from repo failed")
	}
	return &empty.Empty{}, nil
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
