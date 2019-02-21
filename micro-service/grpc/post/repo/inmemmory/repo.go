package inmemmory

import (
	"errors"

	"github.com/flexzuu/benchmark/micro-service/grpc/post/repo/entity"
)

// Repo is used to implement an inmemmory version of  repo.Post
type Repo struct {
	data   map[int64]entity.Post
	nextID int64
}

func (r *Repo) Get(ID int64) (entity.Post, error) {
	p, ok := r.data[ID]
	if !ok {
		return entity.Post{}, errors.New("post not found")
	}
	return p, nil
}
func (r *Repo) Create(AuthorID int64, Headline string, Content string) (entity.Post, error) {
	ID := r.nextID
	r.nextID++

	// check if ID does exist allready

	_, exists := r.data[ID]
	if exists {
		return entity.Post{}, errors.New("no more space") // we ran out of IDs
	}
	p := entity.Post{
		ID,
		AuthorID,
		Headline,
		Content,
	}
	r.data[p.ID] = p
	return p, nil
}
func (r *Repo) Delete(ID int64) error {
	delete(r.data, ID)
	return nil
}

func NewRepo() *Repo {
	data := make(map[int64]entity.Post)
	var nextID int64
	return &Repo{data, nextID}
}
