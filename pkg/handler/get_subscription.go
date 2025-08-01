package handler

import (
	"net/http"

	"github.com/lenkton/effective-mobile-test/pkg/httputil"
	"github.com/lenkton/effective-mobile-test/pkg/middleware"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

type GetSubscription struct {
}

// ServeHTTP implements http.Handler.
// Requires: WithSubscription middleware in chain prior to this
func (h *GetSubscription) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sub := r.Context().Value(middleware.SubscriptionContextKey).(*subscription.Subscription)

	httputil.WriteJSON(w, http.StatusOK, sub)
}
