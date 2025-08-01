package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/lenkton/effective-mobile-test/pkg/configuration"
	"github.com/lenkton/effective-mobile-test/pkg/handler"
	"github.com/lenkton/effective-mobile-test/pkg/middleware"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error: .env file missing")
	}
}

func main() {
	config := configuration.New()

	// WARN: Conn is not thread safe!!!
	db, err := pgx.Connect(context.Background(), config.DatabaseURL)
	if err != nil {
		log.Fatalf("Error: unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	mux := http.NewServeMux()

	subStorage := subscription.NewStorage(db)

	mux.Handle("GET /subscriptions", &handler.ListSubscriptions{Storage: subStorage})
	mux.Handle("GET /subscriptions/{id}", middleware.WithSubscriptionID(middleware.WithSubscription(&handler.GetSubscription{}, subStorage)))
	mux.Handle("POST /subscriptions", &handler.CreateSubscription{Storage: subStorage})
	mux.Handle("DELETE /subscriptions/{id}", middleware.WithSubscriptionID(&handler.DeleteSubscription{Storage: subStorage}))
	mux.Handle("PUT /subscriptions/{id}", middleware.WithSubscriptionID(&handler.UpdateSubscription{DB: db}))

	var handler http.Handler = middleware.NewResultLogger(mux)

	log.Fatal(http.ListenAndServe(config.Address(), handler))
}
