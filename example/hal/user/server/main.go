/*
 * User Service
 *
 * a user service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"

	"github.com/flexzuu/thesis/example/hal/user/api"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	log.Printf("Server started")

	router := api.NewRouter()
	router.Use(cors.Default())
	log.Fatal(router.Run(":4001"))
}