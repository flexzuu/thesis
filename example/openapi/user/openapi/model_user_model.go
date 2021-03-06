/*
 * User Service
 *
 * a user service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UserModel - A User
type UserModel struct {
	Id int64 `json:"id"`

	Email string `json:"email,omitempty"`

	Name string `json:"name,omitempty"`
}
