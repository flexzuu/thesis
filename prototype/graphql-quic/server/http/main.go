package main

import (
	"log"
	"net/http"

	"github.com/flexzuu/thesis/prototype/graphql-quic/graphql"
	"github.com/flexzuu/thesis/prototype/graphql-quic/testdata"
)

func main() {
	certFile, keyFile := testdata.GetCertificatePaths()

	log.Fatal(http.ListenAndServeTLS("localhost:8881", certFile, keyFile, graphql.NewHandler("http")))
}
