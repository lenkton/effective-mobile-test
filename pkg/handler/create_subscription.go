package handler

import (
	"context"
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
	uuid := httputil.NewUUID()
	sub.UserID = uuid
	h.DB.Exec(
		context.Background(),
		"INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5)",
		sub.ServiceName,
		sub.Price,
		sub.UserID,
		sub.StartDate.Time,
		sub.EndDate,
	)
	httputil.WriteJSON(w, http.StatusCreated, sub)
}
