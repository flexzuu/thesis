package main

import (
	"log"
	"net/http"

	"github.com/flexzuu/thesis/graphql/server"
)

func main() {
	log.Fatal(http.ListenAndServeTLS("localhost:8881", "../../../certs/thesis.pem", "../../../certs/thesis-key.pem", server.NewHandler()))
}
