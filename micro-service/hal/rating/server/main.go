/*
 * Post Service
 *
 * a post service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"

	"github.com/flexzuu/benchmark/micro-service/hal/rating/api"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	log.Printf("Server started")

	router := api.NewRouter()
	router.Use(cors.Default())
	log.Fatal(router.Run(":4003"))
}
