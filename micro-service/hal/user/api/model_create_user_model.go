

package api

// CreateUserModel - create user body
type CreateUserModel struct {

	Email string `json:"email,omitempty"`

	Name string `json:"name,omitempty"`
}
