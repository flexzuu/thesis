package repo

import "github.com/flexzuu/benchmark/micro-service/graphql/post/repo/entity"

type Post interface {
	GetById(ID int) (entity.Post, error)
	List() ([]entity.Post, error)
	ListOfAuthor(AuthorID int) ([]entity.Post, error)
	Create(AuthorID int, Headline string, Content string) (entity.Post, error)
	Delete(ID int) error
}
