package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func (b Book) toJSON() []byte {
	t, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return t
}

func FromJson(v []byte) Book {
	book := Book{}

	json.Unmarshal(v, &book)
	return book
}

var books = []Book{
	Book{Title: "Learning spark", Author: "Mukiza", ISBN: "182828282"},
	Book{Title: "Cloud Native go", Author: "Let", ISBN: "887722662"},
	Book{Title: "Learning go lang", Author: "Alex", ISBN: "7262628"},
}

func BookHandlerFunc(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(books)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(res)
}
