package main

import (
	"net/http"

	"log"

	"github.com/go-chi/chi/v5"

)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))

	})

	log.Fatal(http.ListenAndServe(":8000", r))
}
