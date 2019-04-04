package model

import "fmt"

type News struct {
	ID         int
	Content    string
	AuthorID   int
	CommentIDs []int
}

type Author struct {
	ID   int
	Name string
}

type Comment struct {
	ID       int
	AuthorID int
	Content  string
}

var AllNews []News = make([]News, 15)
var AllAuthors []Author = make([]Author, 5)
var AllComments []Comment = make([]Comment, 45)

func init() {
	AllAuthors[0] = Author{
		ID:   0,
		Name: "Hans Peter",
	}
	AllAuthors[1] = Author{
		ID:   1,
		Name: "Olli Molli",
	}
	AllAuthors[2] = Author{
		ID:   2,
		Name: "Spammi One",
	}
	AllAuthors[3] = Author{
		ID:   3,
		Name: "Spammi Two",
	}
	AllAuthors[4] = Author{
		ID:   4,
		Name: "Spammi Three",
	}
	for i, news := range AllNews {
		news.ID = i
		if i < 10 {
			news.AuthorID = 0
		} else {
			news.AuthorID = 1
		}
		ci := 0
		for ci < 3 {
			id := i*3 + ci
			news.CommentIDs = append(news.CommentIDs, id)
			AllComments[id] = Comment{
				AuthorID: ci + 2,
				ID:       id,
				Content:  fmt.Sprint("SPAM"),
			}
			ci = ci + 1
		}
		news.Content = fmt.Sprintf("Nice news #%d", i+1)
		AllNews[i] = news
	}
}
