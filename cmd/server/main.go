package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/lenkton/effective-mobile-test/pkg/configuration"
	"github.com/lenkton/effective-mobile-test/pkg/handler"
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

	mux := http.NewServeMux()

	mux.Handle("GET /subscriptions", &handler.ListSubscriptions{})
	mux.Handle("POST /subscriptions", &handler.CreateSubscription{})

	var handler http.Handler = middleware.NewResultLogger(mux)

	log.Fatal(http.ListenAndServe(config.Address(), handler))
}
