package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/lenkton/effective-mobile-test/pkg/httputil"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

type GetSubscription struct {
	DB *pgx.Conn
}

// ServeHTTP implements http.Handler.
func (h *GetSubscription) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		log.Printf("Error: GetSubscription#ServeHTTP:Atoi: %v\n", err)
		httputil.WriteErrorJSON(w, http.StatusUnprocessableEntity, "malformed subscription id")
		return
	}

	var sub subscription.Subscription
	err = h.DB.QueryRow(
		context.Background(),
		`SELECT id,
		        service_name,
				price,
				user_id,
				start_date,
				end_date
		FROM subscriptions WHERE id = $1`,
		id,
	).Scan(
		&sub.ID,
		&sub.ServiceName,
		&sub.Price,
		&sub.UserID,
		&sub.StartDate,
		&sub.EndDate,
	)
	if err != nil {
		// TODO: use log levels
		log.Printf("Error: GetSubscription#ServeHTTP:QueryRow: %v\n", err)
		httputil.WriteErrorJSON(w, http.StatusNotFound, "subscription not found")
		return
	}

	httputil.WriteJSON(w, http.StatusOK, sub)
}
