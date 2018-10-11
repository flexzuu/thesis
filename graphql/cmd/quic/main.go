package main

import (
	"log"

	"github.com/flexzuu/benchmark/graphql/server"
	"github.com/lucas-clemente/quic-go/h2quic"
)

func main() {
	handler := server.NewHandler()

	log.Fatal(h2quic.ListenAndServeQUIC("localhost:8882", "../../../certs/benchmark.pem", "../../../certs/benchmark-key.pem", handler))
}
