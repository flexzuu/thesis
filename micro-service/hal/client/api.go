// package client

// import (
// 	"encoding/json"
// 	"io/ioutil"

// 	"github.com/flexzuu/benchmark/micro-service/hal/post/api"
// 	"github.com/leibowitz/halgo"
// )

// type Client struct {
// 	url string
// }

// func New(url string) *Client {
// 	return &Client{
// 		url,
// 	}
// }

// func (c *Client) GetById(id int64) (api.PostModel, error) {
// 	r, err := halgo.Navigator(c.url).
// 		Followf("find", halgo.P{"id": id}).
// 		Get()
// 	if err != nil {
// 		return api.PostModel{}, err
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return api.PostModel{}, err
// 	}
// 	var post api.PostModel
// 	json.Unmarshal(body, &post)

// 	return post, nil
// }

// func (c *Client) ListPosts() (api.PostListModel, error) {
// 	r, err := halgo.Navigator(c.url).
// 		Follow("posts").
// 		Get()
// 	if err != nil {
// 		return api.PostListModel{}, err
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return api.PostListModel{}, err
// 	}
// 	var posts api.PostListModel
// 	json.Unmarshal(body, &posts)

// 	return posts, nil
// }

// func (c *Client) ListPostsByAuthor(authorId int64) (api.PostListModel, error) {
// 	r, err := halgo.Navigator(c.url).
// 		Followf("posts", halgo.P{"authorId": authorId}).
// 		Get()
// 	if err != nil {
// 		return api.PostListModel{}, err
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return api.PostListModel{}, err
// 	}
// 	var posts api.PostListModel
// 	json.Unmarshal(body, &posts)

// 	return posts, nil
// }
