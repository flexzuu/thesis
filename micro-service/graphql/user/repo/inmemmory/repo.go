package inmemmory

import (
	"github.com/pkg/errors"

	"github.com/flexzuu/benchmark/micro-service/graphql/user/repo/entity"
)

// Repo is used to implement an inmemmory version of  repo.User
type Repo struct {
	data   map[int]entity.User
	nextID int
}

func (r *Repo) Get(ID int) (entity.User, error) {
	u, ok := r.data[ID]
	if !ok {
		return entity.User{}, errors.New("user not found")
	}
	return u, nil
}
func (r *Repo) Create(Email string, Name string) (entity.User, error) {
	ID := r.nextID
	r.nextID++

	// check if ID does exist allready
	_, exists := r.data[ID]
	if exists {
		return entity.User{}, errors.New("no more space") // we ran out of IDs
	}
	//TODO: check if there is a user with this email already

	u := entity.User{
		ID,
		Email,
		Name,
	}
	err := u.Valid()
	if err != nil {
		return u, errors.Wrap(err, "validation failed")
	}
	r.data[u.ID] = u
	return u, nil
}
func (r *Repo) Delete(ID int) error {
	delete(r.data, ID)
	return nil
}

func NewRepo() *Repo {
	data := make(map[int]entity.User)
	var nextID int
	return &Repo{data, nextID}
}
