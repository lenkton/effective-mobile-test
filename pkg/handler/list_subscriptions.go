package handler

import (
	"log"
	"net/http"

	"github.com/lenkton/effective-mobile-test/pkg/httputil"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

type ListSubscriptions struct {
	Storage *subscription.Storage
}

// ServeHTTP implements http.Handler.
func (h *ListSubscriptions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	subs, err := h.Storage.ListSubscriptions()
	if err != nil {
		httputil.WriteErrorJSON(w, http.StatusInternalServerError, "internal server error")
		log.Printf("Error: ListSubscriptions#ServerHTTP: Storage.ListSubscriptions %v\n", err)
		return
	}

	httputil.WriteJSON(w, http.StatusOK, subs)
}
