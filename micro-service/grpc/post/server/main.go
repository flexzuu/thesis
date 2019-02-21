//go:generate protoc -I ../post --go_out=plugins=grpc:../post ../post/post.proto

package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/flexzuu/benchmark/micro-service/grpc/post/post"
	"github.com/flexzuu/benchmark/micro-service/grpc/post/repo"
	"github.com/flexzuu/benchmark/micro-service/grpc/post/repo/inmemmory"
	"github.com/flexzuu/benchmark/micro-service/grpc/user/user"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement post.PostServiceServer
type server struct {
	postRepo   repo.Post
	userClient user.UserServiceClient
}

// GetPost implements post.PostServiceServer
func (s *server) GetById(ctx context.Context, in *pb.GetPostRequest) (*pb.Post, error) {
	p, err := s.postRepo.Get(in.GetID())
	if err != nil {
		return nil, errors.Wrap(err, "get from repo failed")
	}
	return p.ToProto(), nil
}

// CreatePost implements post.PostServiceServer
func (s *server) Create(ctx context.Context, in *pb.CreatePostRequest) (*pb.Post, error) {
	//Validate AuthorID with user service
	authorID := in.GetAuthorID()
	_, err := s.userClient.GetById(ctx, &user.GetUserRequest{
		ID: authorID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "invalid author id")
	}
	p, err := s.postRepo.Create(authorID, in.GetHeadline(), in.GetContent())
	if err != nil {
		return nil, errors.Wrap(err, "create from repo failed")
	}
	return p.ToProto(), nil
}

// DeletePost implements post.PostServiceServer
func (s *server) Delete(ctx context.Context, in *pb.DeletePostRequest) (*empty.Empty, error) {
	err := s.postRepo.Delete(in.GetID())
	if err != nil {
		return nil, errors.Wrap(err, "delete from repo failed")
	}
	return &empty.Empty{}, nil
}

func main() {
	address := os.Getenv("USER_SERVICE")
	postRepo := inmemmory.NewRepo()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userClient := user.NewUserServiceClient(conn)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &server{
		postRepo,
		userClient,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
