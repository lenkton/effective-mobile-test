package handler

import (
	"log"
	"net/http"

	"github.com/lenkton/effective-mobile-test/pkg/httputil"
	"github.com/lenkton/effective-mobile-test/pkg/middleware"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

type DeleteSubscription struct {
	Storage *subscription.Storage
}

// ServeHTTP implements http.Handler.
// Requires: WithSubscriptionID middleware in chain prior to this
func (h *DeleteSubscription) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.SubscriptionIDContextKey).(int)

	err := h.Storage.DeleteSubscription(id)

	if _, ok := err.(*subscription.ErrorNoRowsAffected); ok {
		log.Printf("Warn: DeleteSubscription#ServeHTTP: DeleteSubscription: %v\n", err)
		httputil.WriteErrorJSON(w, http.StatusNotFound, "subscription not found")
		return
	}
	if err != nil {
		// TODO: use log levels
		log.Printf("Error: DeleteSubscription#ServeHTTP: DeleteSubscription: %v\n", err)
		httputil.WriteErrorJSON(w, http.StatusInternalServerError, "internal server error")
		return
	}

	log.Printf("Info: DeleteSubscription: deleted a record with id %d\n", id)

	w.WriteHeader(http.StatusNoContent)
}
