package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/lenkton/effective-mobile-test/pkg/httputil"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

type contextKey int

// TODO: move this middleware to the subscription service
// so it would become a private const
const (
	SubscriptionContextKey contextKey = iota
	SubscriptionIDContextKey
)

// Requires: WithSubscriptionID middleware in chain prior to this
func WithSubscription(next http.Handler, db *pgx.Conn) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(SubscriptionIDContextKey).(int)

		sub := &subscription.Subscription{}
		err := db.QueryRow(
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
			log.Printf("Error: WithSubscription#ServeHTTP:QueryRow: %v\n", err)
			httputil.WriteErrorJSON(w, http.StatusNotFound, "subscription not found")
			return
		}

		nextContext := context.WithValue(r.Context(), SubscriptionContextKey, sub)
		next.ServeHTTP(w, r.WithContext(nextContext))
	})
}
