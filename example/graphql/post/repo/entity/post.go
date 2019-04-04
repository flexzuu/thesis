package entity

type Post struct {
	ID       int
	AuthorID int
	Headline string
	Content  string //markdown
}
