package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/lenkton/effective-mobile-test/pkg/httputil"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

type ListSubscriptions struct {
	DB *pgx.Conn
}

// ServeHTTP implements http.Handler.
func (h *ListSubscriptions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(context.Background(), "select * from subscriptions")
	if err != nil {
		// TODO: another error code
		httputil.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		log.Printf("Error: ListSubscriptions#ServerHTTP/Query: %v\n", err)
		return
	}
	subs, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[subscription.Subscription])
	if err != nil {
		// TODO: another error code
		httputil.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		log.Printf("Error: ListSubscriptions#ServerHTTP/CollectRows: %v\n", err)
		return
	}

	httputil.WriteJSON(w, http.StatusOK, subs)
}
