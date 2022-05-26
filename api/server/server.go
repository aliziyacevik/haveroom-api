package server

import (
	"net/http"
	"time"
	"log"

	"github.com/alicevvikk/haveroom-api/config"
	"github.com/alicevvikk/haveroom-api/api"

	"github.com/go-chi/chi/v5"

)

type Server struct {
	config		*config.Config
	httpServer	*http.Server
	router		http.Handler
}


func (srv *Server) Run() {
	log.Fatal(srv.httpServer.ListenAndServe())
}

func NewServer(configPath string) *Server {
	conf := config.NewConfig(configPath)
	h := api.NewHandler()
	r := chi.NewRouter()

	r.Get("",    h.Get)
	r.Post("",   h.Post)
	r.Delete("", h.Delete)

	log.Println(conf.Server.Host + ":" + conf.Server.Port)

	httpServer := &http.Server{
		Addr:		conf.Server.Host + ":" + conf.Server.Port,
		ReadTimeout:	time.Duration(conf.Server.ReadTimeout),
		WriteTimeout:	time.Duration(conf.Server.WriteTimeout),
		Handler:	r,
	}


	return &Server{
		config:		conf,
		httpServer:	httpServer,
		router:		r,
	}

}
