package main

import (
	"context"
	"fmt"
	"log"
	"os"

	post "github.com/flexzuu/thesis/micro-service/graphql/post/repo/entity"
	user "github.com/flexzuu/thesis/micro-service/graphql/user/repo/entity"
	"github.com/flexzuu/graphqlt"
)

func main() {
	facadeServiceEndpoint := os.Getenv("FACADE_SERVICE")
	if facadeServiceEndpoint == "" {
		log.Fatalln("please provide FACADE_SERVICE as env var")
	}
	facadeClient := graphqlt.NewClient(facadeServiceEndpoint)

	ListPosts(facadeClient)
	PostDetail(facadeClient, 0)
	AuthorDetail(facadeClient, 0)

}

func ListPosts(facadeClient *graphqlt.Client) {
	// shows post ids+headline
	ctx := context.Background()
	fmt.Println("----------ListPosts----------")
	// fetch posts
	req := graphqlt.NewRequest(`query ListPosts {
		posts: postList {
		  id
		  headline
		}
	  }`)
	var res struct {
		Posts []struct {
			ID       int
			Headline string
		}
	}
	err := facadeClient.Run(ctx, req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("#%d Posts:\n", len(res.Posts))

	for _, post := range res.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.ID)
	}
}

func PostDetail(facadeClient *graphqlt.Client, postID int) {
	fmt.Println("----------PostDetail----------")
	// shows post (headline + content) + authorName and all ratings(avg)
	ctx := context.Background()

	req := graphqlt.NewRequest(`query PostDetail($postID: ID!) {
		post: postGet(id: $postID) {
		  id
		  headline
		  content
		  author {
			id
			name
		  }
		  ratings {
			id
			value
		  }
		}
	  }`)

	req.Var("postID", postID)

	var res struct {
		Post struct {
			post.Post
			Author struct {
				ID   int
				Name string
			}
			Ratings []struct {
				ID    int
				Value int
			}
		}
	}
	err := facadeClient.Run(ctx, req, &res)
	if err != nil {
		log.Fatal(err)
	}

	var avgRating float64

	for _, rating := range res.Post.Ratings {
		avgRating += float64(rating.Value)
	}
	avgRating = avgRating / float64(len(res.Post.Ratings))

	fmt.Printf("%s by %s\nAVG-Rating: %.2f\n%s\n", res.Post.Headline, res.Post.Author.Name, avgRating, res.Post.Content)
}

func AuthorDetail(facadeClient *graphqlt.Client, authorID int) {
	// author name and email
	// shows post ids+headline of author
	// global avg ratings
	fmt.Println("----------AuthorDetail----------")
	ctx := context.Background()
	req := graphqlt.NewRequest(`query AuthorDetail($authorID: ID!) {
		author: userGet(id: $authorID) {
		  id
		  email
		  name
		  posts {
			id
			headline
			ratings {
			  id
			  value
			}
		  }
		}
	  }`)

	req.Var("authorID", authorID)

	var res struct {
		Author struct {
			user.User
			Posts []struct {
				ID       int
				Headline string
				Ratings  []struct {
					ID    int
					Value int
				}
			}
		}
	}
	err := facadeClient.Run(ctx, req, &res)
	if err != nil {
		log.Fatal(err)
	}

	var avgRating float64
	var length int
	for _, post := range res.Author.Posts {

		for _, rating := range post.Ratings {
			avgRating += float64(rating.Value)
		}
		length += len(post.Ratings)
	}
	avgRating = avgRating / float64(length)

	fmt.Printf("%s - %s\n", res.Author.Name, res.Author.Email)
	fmt.Printf("Total AVG-Rating: %.2f\n", avgRating)

	fmt.Printf("#%d Posts:\n", len(res.Author.Posts))

	for _, post := range res.Author.Posts {
		fmt.Printf("\t%s (%d)\n", post.Headline, post.ID)
	}

}
