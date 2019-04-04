package inmemmory

import (
	"errors"

	"github.com/flexzuu/thesis/example/rest/post/repo/entity"
)

// Repo is used to implement an inmemmory version of  repo.Post
type Repo struct {
	data   map[int64]entity.Post
	nextID int64
}

func (r *Repo) GetById(ID int64) (entity.Post, error) {
	p, ok := r.data[ID]
	if !ok {
		return entity.Post{}, errors.New("post not found")
	}
	return p, nil
}
func (r *Repo) List() ([]entity.Post, error) {
	posts := make([]entity.Post, len(r.data))
	for i, post := range r.data {
		posts[i] = post
	}
	return posts, nil
}
func (r *Repo) ListOfAuthor(AuthorID int64) ([]entity.Post, error) {
	posts := make([]entity.Post, 0)
	for _, post := range r.data {
		if post.AuthorID == AuthorID {
			posts = append(posts, post)
		}
	}
	return posts, nil
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
