package main

import (
	"context"
	"fmt"
	"log"
	"os"

	facadeApi "github.com/flexzuu/thesis/example/rest/facade/openapi/client"
)

func main() {
	facadeServiceAddress := os.Getenv("FACADE_SERVICE")
	if facadeServiceAddress == "" {
		log.Fatalln("please provide FACADE_SERVICE as env var")
	}

	facadeCfg := facadeApi.NewConfiguration()
	facadeCfg.BasePath = fmt.Sprintf("http://%s", facadeServiceAddress)
	facadeClient := facadeApi.NewAPIClient(facadeCfg)

	ListPosts(facadeClient.FacadeApi)
	PostDetail(facadeClient.FacadeApi, 0)
	// AuthorDetail(facadeClient, 0)

}

func ListPosts(facadeClient *facadeApi.FacadeApiService) {
	// shows post ids+headline
	ctx := context.Background()
	fmt.Println("----------ListPosts----------")
	// fetch posts
	res, _, err := facadeClient.ListPosts(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("#%d Posts:\n", len(res.Posts))

	for _, post := range res.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.Id)
	}
}

func PostDetail(facadeClient *facadeApi.FacadeApiService, postID int64) {
	fmt.Println("----------PostDetail----------")
	// shows post (headline + content) + authorName and all ratings(avg)
	ctx := context.Background()
	// fetch post by id
	res, _, err := facadeClient.PostDetail(ctx, postID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s by %s\nAVG-Rating: %.2f\n%s\n", res.Post.Headline, res.Author.Name, res.AvgRating, res.Post.Content)
}

func AuthorDetail(facadeClient *facadeApi.FacadeApiService, authorID int64) {
	// author name and email
	// shows post ids+headline of author
	// global avg ratings
	fmt.Println("----------AuthorDetail----------")
	ctx := context.Background()
	res, _, err := facadeClient.AuthorDetail(ctx, authorID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s - %s\n", res.Author.Name, res.Author.Email)
	fmt.Printf("Total AVG-Rating: %.2f\n", res.AvgRating)

	fmt.Printf("#%d Posts:\n", len(res.Posts.Posts))

	for _, post := range res.Posts.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.Id)
	}

}
