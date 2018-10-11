package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/flexzuu/benchmark/common"
	"github.com/gorilla/mux"
)

func NewHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/news", allNewsHandler)
	r.HandleFunc("/news/{id}", singleNewsHandler)
	r.HandleFunc("/authors", allAuthorsHandler)
	r.HandleFunc("/authors/{id}", singleAuthorHandler)
	r.HandleFunc("/comments", allCommentsHandler)
	r.HandleFunc("/comments/{id}", singleCommentHandler)
	return r
}

func allNewsHandler(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(common.AllNews)
	w.Write(json)
}
func singleNewsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json, _ := json.Marshal(common.AllNews[id])
	w.Write(json)
}

func allAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(common.AllAuthors)
	w.Write(json)
}
func singleAuthorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json, _ := json.Marshal(common.AllAuthors[id])
	w.Write(json)
}

func allCommentsHandler(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(common.AllComments)
	w.Write(json)
}
func singleCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json, _ := json.Marshal(common.AllComments[id])
	w.Write(json)
}
