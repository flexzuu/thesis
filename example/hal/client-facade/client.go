package main

import (
	"fmt"
	"log"
	"os"

	api "github.com/flexzuu/thesis/example/hal/facade/api"
	"github.com/leibowitz/halgo"
)

func main() {
	facadeServiceAddress := os.Getenv("FACADE_SERVICE")
	if facadeServiceAddress == "" {
		log.Fatalln("please provide FACADE_SERVICE as env var")
	}
	ListPosts(facadeServiceAddress)
	PostDetail(facadeServiceAddress, 0)
	AuthorDetail(facadeServiceAddress, 0)

}

func ListPosts(facadeServiceAddress string) {
	// shows post ids+headline
	fmt.Println("----------ListPosts----------")
	// fetch posts
	var postsCollection api.PostListModel
	err := halgo.Navigator(facadeServiceAddress).
		Follow("list-posts").
		Unmarshal(&postsCollection)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("#%d Posts:\n", postsCollection.Count)
	for _, post := range postsCollection.Embedded.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.Id)
	}
}

func PostDetail(facadeServiceAddress string, postID int64) {
	fmt.Println("----------PostDetail----------")
	// shows post (headline + content) + authorName and all ratings(avg)
	// fetch post by id
	var postDetail api.PostDetailModel
	err := halgo.Navigator(facadeServiceAddress).
		Followf("post-detail", halgo.P{
			"id": postID,
		}).
		Unmarshal(&postDetail)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s by %s\nAVG-Rating: %.2f\n%s\n", postDetail.Embedded.Post.Headline, postDetail.Embedded.Author.Name, postDetail.AvgRating, postDetail.Embedded.Post.Content)
}

func AuthorDetail(facadeServiceAddress string, authorID int64) {
	// author name and email
	// shows post ids+headline of author
	// global avg ratings
	fmt.Println("----------AuthorDetail----------")
	var authorDetail api.AuthorDetailModel
	err := halgo.Navigator(facadeServiceAddress).
		Followf("author-detail", halgo.P{
			"id": authorID,
		}).
		Unmarshal(&authorDetail)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s - %s\n", authorDetail.Embedded.Author.Name, authorDetail.Embedded.Author.Email)
	fmt.Printf("Total AVG-Rating: %.2f\n", authorDetail.AvgRating)

	fmt.Printf("#%d Posts:\n", len(authorDetail.Embedded.Posts))

	for _, post := range authorDetail.Embedded.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.Id)
	}
}
