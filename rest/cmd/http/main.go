package main

import (
	"log"
	"net/http"

	"github.com/flexzuu/benchmark/rest/server"
)

func main() {
	log.Fatal(http.ListenAndServeTLS("localhost:7771", "../../../certs/benchmark.pem", "../../../certs/benchmark-key.pem", server.NewHandler()))
}
