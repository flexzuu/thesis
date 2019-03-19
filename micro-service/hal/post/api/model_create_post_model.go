

package api

// CreatePostModel - create post body
type CreatePostModel struct {
	AuthorId int64 `json:"authorId"`

	Headline string `json:"headline,omitempty"`

	// content as markdown
	Content string `json:"content,omitempty"`
}
