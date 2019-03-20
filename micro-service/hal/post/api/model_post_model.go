package api

import "github.com/leibowitz/halgo"

// PostModel - A Post
type PostModel struct {
	halgo.Links
	Id int64 `json:"id"`

	Headline string `json:"headline,omitempty"`

	// content as markdown
	Content string `json:"content,omitempty"`
}
