package main

import (
	"log"

	"github.com/flexzuu/thesis/prototype/graphql-quic/graphql"
	"github.com/lucas-clemente/quic-go/h2quic"
)

func main() {
	handler := server.NewHandler()

	log.Fatal(h2quic.ListenAndServeQUIC("localhost:8882", "../../certs/thesis.pem", "../../certs/thesis-key.pem", handler))
}
