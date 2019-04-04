package graphql

import (
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/handler"
	"github.com/flexzuu/thesis/prototype/graphql-quic/model"
	"github.com/gorilla/mux"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct {
	method string
}

func (*query) NewsFeed() []newsResolver {
	news := make([]newsResolver, len(model.AllNews))
	for i, n := range model.AllNews {
		news[i] = newsResolver{
			News: n,
		}
	}
	return news
}

func (q query) Method() string {
	return q.method
}

type newsResolver struct {
	model.News
}

func (n newsResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(n.News.ID))
}
func (n newsResolver) Content() string {
	return n.News.Content
}
func (n newsResolver) Comments() []commentResolver {
	comments := make([]commentResolver, len(n.CommentIDs))
	for i, id := range n.CommentIDs {
		comments[i] = commentResolver{model.AllComments[id]}
	}
	return comments
}

func (n newsResolver) Author() authorResolver {
	author := model.AllAuthors[n.News.AuthorID]
	return authorResolver{
		Author: author,
	}
}

type authorResolver struct {
	model.Author
}

func (a authorResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(a.Author.ID))
}

func (a authorResolver) Name() string {
	return a.Author.Name
}

type commentResolver struct {
	model.Comment
}

func (c commentResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(c.Comment.ID))
}

func (c commentResolver) Content() string {
	return c.Comment.Content
}

func (c commentResolver) Author() authorResolver {
	author := model.AllAuthors[c.Comment.AuthorID]
	return authorResolver{
		Author: author,
	}
}

func NewHandler(method string) http.Handler {
	s := `
	schema {
		query: Query
	}
	type Query {
		newsfeed: [News!]!
		method: String!
	}
	type News {
		id: ID!
		content: String!
		author: Author!
		comments: [Comment!]!
	}
	type Author {
		id: ID!
		name: String!
	}
	type Comment {
		id: ID!
		content: String!
		author: Author!	
	}
`
	schema := graphql.MustParseSchema(s, &query{method})

	r := mux.NewRouter()
	r.Handle("/graphql", &relay.Handler{Schema: schema})
	r.Handle("/playground", handler.Playground("Playground", "/graphql"))
	return r
}
