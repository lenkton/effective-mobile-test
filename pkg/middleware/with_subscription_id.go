package middleware

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/lenkton/effective-mobile-test/pkg/httputil"
)

func WithSubscriptionID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathID := r.PathValue("id")
		id, err := strconv.Atoi(pathID)
		if err != nil {
			log.Printf("Error: WithSubscriptionID#ServeHTTP:Atoi: %v\n", err)
			httputil.WriteErrorJSON(w, http.StatusUnprocessableEntity, "malformed subscription id")
			return
		}

		nextContext := context.WithValue(r.Context(), SubscriptionIDContextKey, id)
		next.ServeHTTP(w, r.WithContext(nextContext))
	})
}
