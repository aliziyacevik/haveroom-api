package main

import (
	"log"

	"github.com/alicevvikk/haveroom-api/api/server"
)

func main() {

	srv := server.NewServer("./config/config.yaml")
	log.Println("listening..")
	srv.Run()

}
