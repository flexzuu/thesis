package inmemmory

import (
	"errors"

	"github.com/flexzuu/thesis/example/graphql/post/repo/entity"
)

// Repo is used to implement an inmemmory version of  repo.Post
type Repo struct {
	data   map[int]entity.Post
	nextID int
}

func (r *Repo) GetById(ID int) (entity.Post, error) {
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
func (r *Repo) ListOfAuthor(AuthorID int) ([]entity.Post, error) {
	posts := make([]entity.Post, 0)
	for _, post := range r.data {
		if post.AuthorID == AuthorID {
			posts = append(posts, post)
		}
	}
	return posts, nil
}
func (r *Repo) Create(AuthorID int, Headline string, Content string) (entity.Post, error) {
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
func (r *Repo) Delete(ID int) error {
	delete(r.data, ID)
	return nil
}

func NewRepo() *Repo {
	data := make(map[int]entity.Post)
	var nextID int
	return &Repo{data, nextID}
}
