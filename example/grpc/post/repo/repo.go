package repo

import "github.com/flexzuu/thesis/example/grpc/post/repo/entity"

type Post interface {
	GetById(ID int64) (entity.Post, error)
	List() ([]entity.Post, error)
	ListOfAuthor(AuthorID int64) ([]entity.Post, error)
	Create(AuthorID int64, Headline string, Content string) (entity.Post, error)
	Delete(ID int64) error
}
