package handler

import (
	"net/http"

	"github.com/lenkton/effective-mobile-test/pkg/httputil"
)

type ListSubscriptions struct{}

// ServeHTTP implements http.Handler.
func (h *ListSubscriptions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httputil.WriteJSON(w, http.StatusOK, []any{})
}
