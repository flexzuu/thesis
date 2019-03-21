package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/flexzuu/benchmark/micro-service/grpc/facade/facade"
	"github.com/flexzuu/benchmark/micro-service/grpc/stats"
	"google.golang.org/grpc"
)

func main() {
	facadeServiceAddress := os.Getenv("FACADE_SERVICE")
	if facadeServiceAddress == "" {
		log.Fatalln("please provide FACADE_SERVICE as env var")
	}

	facadeConn, err := grpc.Dial(facadeServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to facade service: %v", err)
	}
	defer facadeConn.Close()
	facadeClient := facade.NewFacadeServiceClient(facadeConn)
	statsClient := stats.NewStatsClient(facadeConn)

	Reset(statsClient)

	ListPosts(facadeClient)
	PostDetail(facadeClient, 0)
	AuthorDetail(facadeClient, 0)

	Roundtrips(statsClient)
}

func ListPosts(facadeClient facade.FacadeServiceClient) {
	// shows post ids+headline
	ctx := context.Background()
	fmt.Println("----------ListPosts----------")
	// fetch posts
	res, err := facadeClient.ListPosts(ctx, &facade.ListPostsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("#%d Posts:\n", len(res.Posts))

	for _, post := range res.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.ID)
	}
}

func PostDetail(facadeClient facade.FacadeServiceClient, postID int64) {
	fmt.Println("----------PostDetail----------")
	// shows post (headline + content) + authorName and all ratings(avg)
	ctx := context.Background()
	// fetch post by id
	res, err := facadeClient.PostDetail(ctx, &facade.PostDetailRequest{
		ID: postID,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s by %s\nAVG-Rating: %.2f\n%s\n", res.Post.Headline, res.Author.Name, res.AvgRating, res.Post.Content)
}

func AuthorDetail(facadeClient facade.FacadeServiceClient, authorID int64) {
	// author name and email
	// shows post ids+headline of author
	// global avg ratings
	fmt.Println("----------AuthorDetail----------")
	ctx := context.Background()
	res, err := facadeClient.AuthorDetail(ctx, &facade.AuthorDetailRequest{
		ID: authorID,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s - %s\n", res.Author.Name, res.Author.Email)
	fmt.Printf("Total AVG-Rating: %.2f\n", res.AvgRating)

	fmt.Printf("#%d Posts:\n", len(res.Posts))

	for _, post := range res.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.ID)
	}

}

func Roundtrips(statsClient stats.StatsClient) {
	// shows post ids+headline
	ctx := context.Background()
	rt, err := statsClient.RoundTrips(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Roundtrips to facade: %d\n", rt.Count)

}

func Reset(statsClient stats.StatsClient) {
	// shows post ids+headline
	ctx := context.Background()
	_, err := statsClient.Reset(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalln(err)
	}
}
