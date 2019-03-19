

package api

// PostModel - A Post
type PostModel struct {
	Id int64 `json:"id"`

	AuthorId int64 `json:"authorId"`

	Headline string `json:"headline,omitempty"`

	// content as markdown
	Content string `json:"content,omitempty"`
}
