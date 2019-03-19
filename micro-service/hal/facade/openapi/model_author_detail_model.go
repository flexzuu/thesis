/*
 * Facade Service
 *
 * a facade service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	post "github.com/flexzuu/benchmark/micro-service/rest/post/openapi"
	user "github.com/flexzuu/benchmark/micro-service/rest/user/openapi"
)

// AuthorDetailModel - Author with more info included
type AuthorDetailModel struct {
	Author user.UserModel `json:"author"`

	Posts post.PostListModel `json:"posts"`

	AvgRating float64 `json:"avgRating"`
}