package repo

import "github.com/flexzuu/benchmark/micro-service/grpc/post/repo/entity"

type Post interface {
	Get(ID int64) (entity.Post, error)
	Create(AuthorID int64, Headline string, Content string) (entity.Post, error)
	Delete(ID int64) error
}
