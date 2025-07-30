package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/lenkton/effective-mobile-test/pkg/httputil"
)

type DeleteSubscription struct {
	DB *pgx.Conn
}

// ServeHTTP implements http.Handler.
func (h *DeleteSubscription) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		log.Printf("Error: DeleteSubscription#ServeHTTP:Atoi: %v\n", err)
		httputil.WriteErrorJSON(w, http.StatusUnprocessableEntity, "malformed subscription id")
		return
	}

	commandTag, err := h.DB.Exec(
		context.Background(),
		`DELETE FROM subscriptions WHERE id = $1`,
		id,
	)
	if err != nil {
		// TODO: use log levels
		log.Printf("Error: DeleteSubscription#ServeHTTP:Exec: %v\n", err)
		httputil.WriteErrorJSON(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if commandTag.RowsAffected() != 1 {
		log.Printf("Warn: DeleteSubscription#ServeHTTP:Exec: affected %v rows\n", commandTag.RowsAffected())
		httputil.WriteErrorJSON(w, http.StatusNotFound, "subscription not found")
		return
	}
	log.Printf("Info: DeleteSubscription: deleted %d records for id %d\n", commandTag.RowsAffected(), id)

	w.WriteHeader(http.StatusNoContent)
}
