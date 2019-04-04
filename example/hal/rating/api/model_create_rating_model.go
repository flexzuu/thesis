

package api

// CreateRatingModel - create rating body
type CreateRatingModel struct {
	PostId int64 `json:"postId"`

	Rating int32 `json:"rating,omitempty"`
}
