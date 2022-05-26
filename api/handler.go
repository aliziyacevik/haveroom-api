package api

import (
	"net/http"
	"log"
)

type Handler interface{
	Get(http.ResponseWriter,    *http.Request)
	Post(http.ResponseWriter,   *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type handler struct {}

func NewHandler() Handler{
	return new(handler)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("get")
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	log.Println("post")
}

func (h *handler) Delete(w http.ResponseWriter, r*http.Request) {
	log.Println("delete")
}



