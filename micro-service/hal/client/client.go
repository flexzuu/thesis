package main

import (
	"fmt"
	"log"
	"os"

	post "github.com/flexzuu/thesis/micro-service/hal/post/api"
	rating "github.com/flexzuu/thesis/micro-service/hal/rating/api"
	user "github.com/flexzuu/thesis/micro-service/hal/user/api"
	"github.com/leibowitz/halgo"
)

func main() {
	postServiceAddress := os.Getenv("POST_SERVICE")
	if postServiceAddress == "" {
		log.Fatalln("please provide POST_SERVICE as env var")
	}

	userServiceAddress := os.Getenv("USER_SERVICE")
	if userServiceAddress == "" {
		log.Fatalln("please provide USER_SERVICE as env var")
	}

	ratingServiceAddress := os.Getenv("RATING_SERVICE")
	if ratingServiceAddress == "" {
		log.Fatalln("please provide RATING_SERVICE as env var")
	}

	ListPosts(postServiceAddress)
	PostDetail(postServiceAddress, 0)
	AuthorDetail(userServiceAddress, ratingServiceAddress, 0)
}

func ListPosts(postServiceAddress string) {
	// shows post ids+headline
	fmt.Println("----------ListPosts----------")
	// fetch posts
	var postsCollection post.PostListModel
	err := halgo.Navigator(postServiceAddress).
		Follow("posts").
		Unmarshal(&postsCollection)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("#%d Posts:\n", postsCollection.Count)
	for _, post := range postsCollection.Embedded.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.Id)
	}
}

func PostDetail(postServiceAddress string, postID int64) {
	fmt.Println("----------PostDetail----------")
	// shows post (headline + content) + authorName and all ratings(avg)
	// fetch post by id
	var post post.PostModel
	err := halgo.Navigator(postServiceAddress).
		Followf("find", halgo.P{
			"id": postID,
		}).
		Unmarshal(&post)
	if err != nil {
		log.Fatal(err)
	}
	var author user.UserModel
	err = halgo.Navigator(postServiceAddress).
		Followf("find", halgo.P{
			"id": postID,
		}).
		Follow("author").
		Unmarshal(&author)
	if err != nil {
		log.Fatal(err)
	}
	var ratings rating.RatingListModel
	err = halgo.Navigator(postServiceAddress).
		Followf("find", halgo.P{
			"id": postID,
		}).
		Follow("ratings").
		Unmarshal(&ratings)
	if err != nil {
		log.Fatal(err)
	}
	var avgRating float64

	for _, rating := range ratings.Embedded.Ratings {
		avgRating += float64(rating.Value)
	}
	avgRating = avgRating / float64(ratings.Count)

	fmt.Printf("%s by %s\nAVG-Rating: %.2f\n%s\n", post.Headline, author.Name, avgRating, post.Content)
}

func AuthorDetail(userServiceAddress, ratingServiceAddress string, authorID int64) {
	// author name and email
	// shows post ids+headline of author
	// global avg ratings
	fmt.Println("----------AuthorDetail----------")
	var author user.UserModel
	err := halgo.Navigator(userServiceAddress).
		Followf("find", halgo.P{
			"id": authorID,
		}).
		Unmarshal(&author)
	if err != nil {
		log.Fatal(err)
	}
	var postsCollection post.PostListModel
	err = halgo.Navigator(userServiceAddress).
		Followf("find", halgo.P{
			"id": authorID,
		}).
		Follow("posts").
		Unmarshal(&postsCollection)
	if err != nil {
		log.Fatal(err)
	}
	var avgRating float64
	var length int
	posts := postsCollection.Embedded.Posts

	for _, p := range posts {
		var ratings rating.RatingListModel
		err = halgo.Navigator(ratingServiceAddress).
			Followf("ratings", halgo.P{
				"postId": p.Id,
			}).
			Unmarshal(&ratings)
		if err != nil {
			log.Fatal(err)
		}

		for _, rating := range ratings.Embedded.Ratings {
			avgRating += float64(rating.Value)
		}
		length += ratings.Count
	}
	avgRating = avgRating / float64(length)
	fmt.Printf("%s - %s\n", author.Name, author.Email)
	fmt.Printf("Total AVG-Rating: %.2f\n", avgRating)

	fmt.Printf("#%d Posts:\n", len(posts))

	for _, post := range posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.Id)
	}
}
