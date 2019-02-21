package inmemmory

import (
	"errors"

	"github.com/flexzuu/benchmark/micro-service/grpc/user/repo/entity"
)

// Repo is used to implement an inmemmory version of  repo.User
type Repo struct {
	data   map[int64]entity.User
	nextID int64
}

func (r *Repo) Get(ID int64) (entity.User, error) {
	p, ok := r.data[ID]
	if !ok {
		return entity.User{}, errors.New("user not found")
	}
	return p, nil
}
func (r *Repo) Create(Email string, Name string) (entity.User, error) {
	ID := r.nextID
	r.nextID++

	// check if ID does exist allready

	_, exists := r.data[ID]
	if exists {
		return entity.User{}, errors.New("no more space") // we ran out of IDs
	}
	p := entity.User{
		ID,
		Email,
		Name,
	}
	r.data[p.ID] = p
	return p, nil
}
func (r *Repo) Delete(ID int64) error {
	delete(r.data, ID)
	return nil
}

func NewRepo() *Repo {
	data := make(map[int64]entity.User)
	var nextID int64
	return &Repo{data, nextID}
}
