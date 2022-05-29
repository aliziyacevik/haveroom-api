package http

import (
	"log"	
	"net/http"

	"github.com/go-chi/chi/v5"

)

type Api struct {
	Router 		*chi.Mux
}

func NewApi() *Api {
	r := chi.NewRouter()
	api := &Api{
		Router:		r,
	}
	
	r.Get("/", api.Get)
	r.Post("/", api.Post)
	
	return api

}

func (a *Api) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("get")
}

func (a *Api) Post(w http.ResponseWriter, r *http.Request) {
	log.Println("")
}
