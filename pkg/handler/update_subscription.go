package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lenkton/effective-mobile-test/pkg/httputil"
	"github.com/lenkton/effective-mobile-test/pkg/middleware"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

type UpdateSubscription struct {
	Storage *subscription.Storage
}

// ServeHTTP implements http.Handler.
// Requires: WithSubscriptionID middleware in chain prior to this
func (h *UpdateSubscription) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.SubscriptionIDContextKey).(int)

	sub := &subscription.Subscription{}
	err := json.NewDecoder(r.Body).Decode(&sub)
	if err != nil {
		httputil.WriteErrorJSON(w, http.StatusUnprocessableEntity, "malformed request body")
		log.Printf("Error: UpdateSubscription#ServeHTTP:Decode: %v\n", err)
		return
	}

	sub, err = h.Storage.UpdateSubscription(id, sub)

	if _, ok := err.(*subscription.ErrorSubscriptionNotFound); ok {
		log.Printf("Warn: UpdateSubscription#ServeHTTP: UpdateSubscription: %v\n", err)
		httputil.WriteErrorJSON(w, http.StatusNotFound, "subscription not found")
		return
	}
	if err != nil {
		// TODO: use log levels
		log.Printf("Error: UpdateSubscription#ServeHTTP: UpdateSubscription: %v\n", err)
		httputil.WriteErrorJSON(w, http.StatusInternalServerError, "internal server error")
		return
	}

	log.Printf("Info: UpdateSubscription#ServeHTTP: updated subscription into %v\n", sub)

	httputil.WriteJSON(w, http.StatusOK, sub)
}
