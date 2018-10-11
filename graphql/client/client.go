package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func FetchNewsFeed(baseURL string, t time.Duration, c http.Client) {
	time.Sleep(t)
	s := `{"operationName":null,"variables":{},"query":"{\n  newsfeed {\n    id\n    content\n  }\n}\n"}`
	resp, err := c.Post(baseURL+"/graphql", "application/json", strings.NewReader(s))
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	_ = body
}

func FetchNewsFeedWithAuthor(baseURL string, t time.Duration, c http.Client) {
	time.Sleep(t)
	s := `{"operationName":null,"variables":{},"query":"{\n  newsfeed {\n    id\n    content\n    author {\n      id\n      name\n    }\n  }\n}\n"}`
	resp, err := c.Post(baseURL+"/graphql", "application/json", strings.NewReader(s))
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	_ = body
}

func FetchNewsFeedWithAuthorAndComments(baseURL string, t time.Duration, c http.Client) {
	time.Sleep(t)
	s := `{"operationName":null,"variables":{},"query":"{\n  newsfeed {\n    id\n    content\n    author {\n      id\n      name\n    }\n    comments {\n      id\n      content\n      author {\n        id\n        name\n      }\n    }\n  }\n}\n"}`
	resp, err := c.Post(baseURL+"/graphql", "application/json", strings.NewReader(s))
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	_ = body
}
