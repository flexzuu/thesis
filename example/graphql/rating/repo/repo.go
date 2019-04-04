package repo

import "github.com/flexzuu/thesis/example/graphql/rating/repo/entity"

type Rating interface {
	GetById(ID int) (entity.Rating, error)
	ListOfPost(PostID int) ([]entity.Rating, error)
	Create(PostID int, value int) (entity.Rating, error)
	Delete(ID int) error
}
