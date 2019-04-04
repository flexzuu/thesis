package main

import (
	"log"
	"net/http"

	"github.com/flexzuu/thesis/rest/server"
)

func main() {
	log.Fatal(http.ListenAndServeTLS("localhost:7771", "../../../certs/thesis.pem", "../../../certs/thesis-key.pem", server.NewHandler()))
}
