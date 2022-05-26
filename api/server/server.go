package server

import (
	"net/http"
	"time"
	"log"
	"github.com/alicevvikk/haveroom-api/config"
)

type Server struct {
	Config		*config.Config
	httpServer 	*http.Server	

}

func (srv *Server) Run() {
	log.Fatal(srv.httpServer.ListenAndServe())
}

func NewServer(configPath string) *Server {
	conf := config.NewConfig(configPath)

	httpServer := &http.Server{
		Addr:		conf.Server.Host + ":" + conf.Server.Port,
		ReadTimeout:	time.Duration(conf.Server.ReadTimeout),
		WriteTimeout:	time.Duration(conf.Server.WriteTimeout),
	}
	

	return &Server{
		Config:		conf,
		httpServer:	httpServer,
	}
		
}
