package story

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("tmpl/index.html"))

type handler struct {
	s Story
}

func NewHandler(s Story) handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	if chapter, ok := h.s[path[1:]]; ok {
		err := tmpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Chapter not found.", http.StatusNotFound)
	}
}
