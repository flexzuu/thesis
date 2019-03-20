package repo

import "github.com/flexzuu/benchmark/micro-service/hal/user/repo/entity"

type User interface {
	Get(ID int64) (entity.User, error)
	Create(Email string, Name string) (entity.User, error)
	Delete(ID int64) error
}
