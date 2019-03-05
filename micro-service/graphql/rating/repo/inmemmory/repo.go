package inmemmory

import (
	"errors"

	"github.com/flexzuu/benchmark/micro-service/graphql/rating/repo/entity"
)

// Repo is used to implement an inmemmory version of  repo.Rating
type Repo struct {
	data   map[int]entity.Rating
	nextID int
}

func (r *Repo) GetById(ID int) (entity.Rating, error) {
	rating, ok := r.data[ID]
	if !ok {
		return entity.Rating{}, errors.New("rating not found")
	}
	return rating, nil
}

func (r *Repo) ListOfPost(PostID int) ([]entity.Rating, error) {
	ratings := make([]entity.Rating, 0)
	for _, rating := range r.data {
		if rating.PostID == PostID {
			ratings = append(ratings, rating)
		}
	}
	return ratings, nil
}
func (r *Repo) Create(PostID int, value int) (entity.Rating, error) {
	ID := r.nextID
	r.nextID++

	// check if ID does exist allready

	_, exists := r.data[ID]
	if exists {
		return entity.Rating{}, errors.New("no more space") // we ran out of IDs
	}
	Value := value
	rating := entity.Rating{
		ID,
		PostID,
		Value,
	}
	r.data[rating.ID] = rating
	return rating, nil
}
func (r *Repo) Delete(ID int) error {
	delete(r.data, ID)
	return nil
}

func NewRepo() *Repo {
	data := make(map[int]entity.Rating)
	var nextID int
	return &Repo{data, nextID}
}
