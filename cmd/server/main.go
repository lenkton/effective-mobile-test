package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/lenkton/effective-mobile-test/pkg/configuration"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error: .env file missing")
	}
}

func main() {
	config := configuration.New()

	log.Fatal(http.ListenAndServe(config.Address(), nil))
}
