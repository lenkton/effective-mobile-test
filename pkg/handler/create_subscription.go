package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/lenkton/effective-mobile-test/pkg/httputil"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

type CreateSubscription struct {
	DB *pgx.Conn
}

// ServeHTTP implements http.Handler.
func (h *CreateSubscription) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sub := &subscription.Subscription{}
	err := json.NewDecoder(r.Body).Decode(&sub)
	if err != nil {
		httputil.WriteErrorJSON(w, http.StatusUnprocessableEntity, "malformed request body")
		log.Printf("Error: CreateSubscription#ServeHTTP:Decode: %v\n", err)
		return
	}

	err = h.DB.QueryRow(
		context.Background(),
		`INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		sub.ServiceName,
		sub.Price,
		sub.UserID,
		sub.StartDate.Time,
		sub.EndDate,
	).Scan(&sub.ID)
	if err != nil {
		httputil.WriteErrorJSON(w, http.StatusInternalServerError, "error saving subscription")
		log.Printf("Error: CreateSubscription#ServeHTTP:QueryRow: %v\n", err)
		return
	}
	httputil.WriteJSON(w, http.StatusCreated, sub)
}
