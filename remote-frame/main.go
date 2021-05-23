package main

import (
	"net/http"

	"github.com/a-h/go-hotwire-todo/remote-frame/templates"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Turbo-Frame"},
	})
	http.Handle("/remote-frame", c.Handler(IndexHandler{}))
	http.ListenAndServe(":8001", nil)
}

type IndexHandler struct{}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	templates.Remote().Render(r.Context(), w)
}
