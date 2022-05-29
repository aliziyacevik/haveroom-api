package http

import (
	"net/http"
	"log"
//	"time"

	"github.com/alicevvikk/haveroom-api/config"

)

type Server struct {
	config	   		*config.Config 
	httpServer 		*http.Server
}

func (srv *Server) Run() {
	log.Printf("Listening on %s..", srv.config.Server.Port)
	log.Fatal(srv.httpServer.ListenAndServe())
}


func NewServer(configPath string) *Server {
	cfg := config.NewConfig(configPath)
	api := NewApi()
	
	httpServer := &http.Server{
		Addr:		cfg.Server.Port,
//		ReadTimeout:	time.Duration(cfg.Server.ReadTimeout),
//		WriteTimeout:	time.Duration(cfg.Server.WriteTimeout),
		Handler:	api.Router,
	}

	
	return &Server{
		config:		cfg,
		httpServer:	httpServer,
	}

}
