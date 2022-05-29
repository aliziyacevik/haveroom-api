package main

import (
	"github.com/alicevvikk/haveroom-api/internal/http"
)

func main() {
	server := http.NewServer("./config/config.yaml")
	server.Run()
	
}
