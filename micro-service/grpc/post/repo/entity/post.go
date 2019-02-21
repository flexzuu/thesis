package entity

type Post struct {
	ID       int64
	AuthorID int64
	Headline string
	Content  string //markdown
}
