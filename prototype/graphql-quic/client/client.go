package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/flexzuu/thesis/prototype/graphql-quic/testdata"
	"github.com/lucas-clemente/quic-go/h2quic"
)

var operation = `{
	"operationName":null,
	"variables":{},
	"query":"{\n method \n newsfeed {\n    id\n    content\n    author {\n      id\n      name\n    }\n    comments {\n      id\n      content\n      author {\n        id\n        name\n      }\n    }\n  }\n}\n"
}`

func main() {
	method := flag.String("method", "http", "http|quic")
	flag.Parse()
	switch *method {
	case "http":
		FetchHTTP()
	case "quic":
		FetchQUIC()
	}
}
func FetchQUIC() {
	roundTripper := &h2quic.RoundTripper{
		TLSClientConfig: &tls.Config{
			RootCAs: testdata.GetRootCA(),
		},
	}
	defer roundTripper.Close()
	c := &http.Client{
		Transport: roundTripper,
	}
	resp, err := c.Post("https://localhost:8882/graphql", "application/json", strings.NewReader(operation))
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	fmt.Print(string(body))
}

func FetchHTTP() {
	trans := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: testdata.GetRootCA(),
		},
	}
	c := &http.Client{
		Transport: trans,
	}
	resp, err := c.Post("https://localhost:8881/graphql", "application/json", strings.NewReader(operation))
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error getting news %v", err)
	}
	fmt.Print(string(body))
}
