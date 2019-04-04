package main

import (
	"log"

	"github.com/flexzuu/thesis/prototype/graphql-quic/graphql"
	"github.com/flexzuu/thesis/prototype/graphql-quic/testdata"
	"github.com/lucas-clemente/quic-go/h2quic"
)

func main() {
	handler := graphql.NewHandler("quic")
	certFile, keyFile := testdata.GetCertificatePaths()
	log.Fatal(h2quic.ListenAndServeQUIC("localhost:8882", certFile, keyFile, handler))
}
