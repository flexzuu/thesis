package main

import (
	"log"

	"github.com/flexzuu/benchmark/rest/server"
	"github.com/lucas-clemente/quic-go/h2quic"
)

func main() {
	handler := server.NewHandler()

	log.Fatal(h2quic.ListenAndServeQUIC("localhost:7772", "../../../certs/benchmark.pem", "../../../certs/benchmark-key.pem", handler))
}
