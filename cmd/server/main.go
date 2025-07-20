package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/lenkton/effective-mobile-test/pkg/configuration"
	"github.com/lenkton/effective-mobile-test/pkg/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error: .env file missing")
	}
}

func main() {
	config := configuration.New()

	var mux http.Handler = http.NewServeMux()

	mux = middleware.NewResultLogger(mux)

	log.Fatal(http.ListenAndServe(config.Address(), mux))
}
