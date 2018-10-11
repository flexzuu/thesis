package main

import (
	"log"
	"net/http"

	"github.com/flexzuu/benchmark/graphql/server"
)

func main() {
	log.Fatal(http.ListenAndServeTLS("localhost:8881", "../../../certs/benchmark.pem", "../../../certs/benchmark-key.pem", server.NewHandler()))
}
