package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/flexzuu/benchmark/common"
)

func FetchNewsFeed(baseURL string, t time.Duration, c http.Client) {
	time.Sleep(t)
	resp, err := c.Get(baseURL + "/news")
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	var AllNews []common.News
	err = json.Unmarshal(body, &AllNews)
	if err != nil {
		log.Fatalf("cant unmarshal %v", err)
	}
	// for _, news := range AllNews {
	// 	fmt.Println(news.Content)
	// }
}

func FetchNewsFeedWithAuthor(baseURL string, t time.Duration, c http.Client) {
	time.Sleep(t)
	resp, err := c.Get(baseURL + "/news")
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	var AllNews []common.News
	err = json.Unmarshal(body, &AllNews)
	if err != nil {
		log.Fatalf("can't unmarshal %v", err)
	}
	for _, news := range AllNews {
		time.Sleep(t)
		resp, err := c.Get(fmt.Sprintf(baseURL+"/authors/%d", news.AuthorID))
		if err != nil {
			log.Fatalf("error getting author %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("error getting author %v", err)
		}
		var author common.Author
		err = json.Unmarshal(body, &author)
		if err != nil {
			log.Fatalf("can't unmarshal %v", err)
		}
		// fmt.Printf("%s by %s\n", news.Content, author.Name)
	}
}

func FetchNewsFeedWithAuthorAndComments(baseURL string, t time.Duration, c http.Client) {
	time.Sleep(t)
	resp, err := c.Get(baseURL + "/news")
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	var AllNews []common.News
	err = json.Unmarshal(body, &AllNews)
	if err != nil {
		log.Fatalf("can't unmarshal %v", err)
	}
	for _, news := range AllNews {
		time.Sleep(t)
		resp, err := c.Get(fmt.Sprintf(baseURL+"/authors/%d", news.AuthorID))
		if err != nil {
			log.Fatalf("error getting author %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("error getting author %v", err)
		}
		var author common.Author
		err = json.Unmarshal(body, &author)
		if err != nil {
			log.Fatalf("can't unmarshal %v", err)
		}
		var comments = make([]common.Comment, len(news.CommentIDs))
		for i, commentID := range news.CommentIDs {
			time.Sleep(t)
			resp, err := c.Get(fmt.Sprintf(baseURL+"/comments/%d", commentID))
			if err != nil {
				log.Fatalf("error getting comment %v", err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("error getting comment %v", err)
			}
			err = json.Unmarshal(body, &comments[i])
			if err != nil {
				log.Fatalf("can't unmarshal %v", err)
			}
		}
		var authors = make([]common.Author, len(comments))
		for i, comment := range comments {
			time.Sleep(t)
			resp, err := c.Get(fmt.Sprintf(baseURL+"/authors/%d", comment.AuthorID))
			if err != nil {
				log.Fatalf("error getting author %v", err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("error getting author %v", err)
			}
			err = json.Unmarshal(body, &authors[i])
			if err != nil {
				log.Fatalf("can't unmarshal %v", err)
			}
		}
		// fmt.Printf("%s by %s\n", news.Content, author.Name)
		// for i, c := range comments {
		// 	fmt.Printf("\t%s by %s\n", c.Content, authors[i].Name)
		// }
	}
}
