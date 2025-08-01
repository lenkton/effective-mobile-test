package middleware

import (
	"context"
	"log"
	"net/http"

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
func WithSubscription(next http.Handler, storage *subscription.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(SubscriptionIDContextKey).(int)

		sub, err := storage.GetSubscription(id)

		if _, ok := err.(*subscription.ErrorSubscriptionNotFound); ok {
			log.Printf("Warn: WithSubscription#ServeHTTP:QueryRow: %v\n", err)
			httputil.WriteErrorJSON(w, http.StatusNotFound, "subscription not found")
			return
		}
		if err != nil {
			// TODO: use log levels
			log.Printf("Error: WithSubscription#ServeHTTP:QueryRow: %v\n", err)
			httputil.WriteErrorJSON(w, http.StatusInternalServerError, "internal server error")
			return
		}

		nextContext := context.WithValue(r.Context(), SubscriptionContextKey, sub)
		next.ServeHTTP(w, r.WithContext(nextContext))
	})
}
