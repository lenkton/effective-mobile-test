package handler

import (
	"net/http"

	"github.com/lenkton/effective-mobile-test/pkg/httputil"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

type CreateSubscription struct{}

// ServeHTTP implements http.Handler.
func (h *CreateSubscription) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sub := &subscription.Subscription{}
	httputil.WriteJSON(w, http.StatusCreated, sub)
}
