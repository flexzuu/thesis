package main

import (
	"log"

	"github.com/flexzuu/thesis/example/hal/facade/api"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	log.Printf("Server started")

	router := api.NewRouter()
	router.Use(cors.Default())
	log.Fatal(router.Run(":4000"))
}
