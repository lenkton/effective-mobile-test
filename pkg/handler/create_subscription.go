package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lenkton/effective-mobile-test/pkg/httputil"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

type CreateSubscription struct {
	Storage *subscription.Storage
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

	id, err := h.Storage.CreateSubscription(sub)

	if err != nil {
		httputil.WriteErrorJSON(w, http.StatusInternalServerError, "error saving subscription")
		log.Printf("Error: CreateSubscription#ServeHTTP:QueryRow: %v\n", err)
		return
	}
	sub.ID = id
	httputil.WriteJSON(w, http.StatusCreated, sub)
}
