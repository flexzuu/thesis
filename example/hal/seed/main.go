package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"

	post "github.com/flexzuu/thesis/example/hal/post/api"
	rating "github.com/flexzuu/thesis/example/hal/rating/api"
	user "github.com/flexzuu/thesis/example/hal/user/api"
	"github.com/leibowitz/halgo"
)

// seed the services with some test data
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

	var createUser = func(usr user.CreateUserModel) user.UserModel {
		b, err := json.Marshal(usr)

		if err != nil {
			log.Fatal(err)
		}
		r, err := halgo.Navigator(userServiceAddress).
			Follow("users").
			Post("application/json", bytes.NewBuffer(b))
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		var user user.UserModel
		err = json.Unmarshal(body, &user)
		if err != nil {
			log.Fatal(err)
		}
		return user
	}

	//createTestUser
	john := createUser(user.CreateUserModel{
		Email: "john@example.com",
		Name:  "John Doe",
	})
	//createTestUser
	jane := createUser(user.CreateUserModel{
		Email: "jane@example.com",
		Name:  "Jane Doe",
	})

	// create fake post
	var createPost = func(author user.UserModel, headline string) post.PostModel {
		b, err := json.Marshal(post.CreatePostModel{
			AuthorId: author.Id,
			Headline: headline,
			Content: `Markdown is a lightweight markup language based on the formatting conventions that people naturally use in email.  As [John Gruber] writes on the [Markdown site][df1]
	
	> The overriding design goal for Markdown's
	> formatting syntax is to make it as readable
	> as possible. The idea is that a
	> Markdown-formatted document should be
	> publishable as-is, as plain text, without
	> looking like it's been marked up with tags
	> or formatting instructions.
	
	This text you see here is *actually* written in Markdown! To get a feel for Markdown's syntax, type some text into the left window and watch the results in the right.`,
		})

		if err != nil {
			log.Fatal(err)
		}
		r, err := halgo.Navigator(postServiceAddress).
			Follow("posts").
			Post("application/json", bytes.NewBuffer(b))
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		var post post.PostModel
		err = json.Unmarshal(body, &post)
		if err != nil {
			log.Fatal(err)
		}
		return post
	}

	// create fake ratings
	var rate = func(postID int64) {

		i := 1
		for i <= 5 /*num of fake reviews*/ {
			rating := rating.CreateRatingModel{
				PostId: postID,
				Rating: rand.Int31n(5),
			}
			b, err := json.Marshal(rating)

			if err != nil {
				log.Fatal(err)
			}
			_, err = halgo.Navigator(ratingServiceAddress).
				Follow("ratings").
				Post("application/json", bytes.NewBuffer(b))
			if err != nil {
				log.Fatal(err)
			}

			i++
		}
	}

	worstBlogPost1 := createPost(john, "Worst Blog Post 1")
	worstBlogPost2 := createPost(john, "Worst Blog Post 2")
	worstBlogPost3 := createPost(john, "Worst Blog Post 3")
	bestBlogPost1 := createPost(jane, "Best Blog Post 1")
	bestBlogPost2 := createPost(jane, "Best Blog Post 2")
	bestBlogPost3 := createPost(jane, "Best Blog Post 3")

	rate(worstBlogPost1.Id)
	rate(worstBlogPost2.Id)
	rate(worstBlogPost3.Id)
	rate(bestBlogPost1.Id)
	rate(bestBlogPost2.Id)
	rate(bestBlogPost3.Id)

}
