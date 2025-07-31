package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/lenkton/effective-mobile-test/pkg/httputil"
	"github.com/lenkton/effective-mobile-test/pkg/subscription"
)

type UpdateSubscription struct {
	DB *pgx.Conn
}

// ServeHTTP implements http.Handler.
func (h *UpdateSubscription) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		log.Printf("Error: UpdateSubscription#ServeHTTP:Atoi: %v\n", err)
		httputil.WriteErrorJSON(w, http.StatusUnprocessableEntity, "malformed subscription id")
		return
	}

	sub := &subscription.Subscription{}
	err = json.NewDecoder(r.Body).Decode(&sub)
	if err != nil {
		httputil.WriteErrorJSON(w, http.StatusUnprocessableEntity, "malformed request body")
		log.Printf("Error: UpdateSubscription#ServeHTTP:Decode: %v\n", err)
		return
	}

	err = h.DB.QueryRow(
		context.Background(),
		`UPDATE subscriptions
		SET service_name=$2,
		    price=$3,
			user_id=$4,
			start_date=$5,
			end_date=$6
		WHERE id = $1
		RETURNING id,
		          service_name,
				  price,
				  user_id,
				  start_date,
				  end_date`,
		id,
		sub.ServiceName,
		sub.Price,
		sub.UserID,
		sub.StartDate,
		sub.EndDate,
	).Scan(
		&sub.ID,
		&sub.ServiceName,
		&sub.Price,
		&sub.UserID,
		&sub.StartDate,
		&sub.EndDate,
	)
	if err == pgx.ErrNoRows {
		log.Printf("Warn: UpdateSubscription#ServeHTTP:Exec: %v\n", err)
		httputil.WriteErrorJSON(w, http.StatusNotFound, "subscription not found")
		return
	}
	if err != nil {
		// TODO: use log levels
		log.Printf("Error: UpdateSubscription#ServeHTTP:Exec: %v\n", err)
		httputil.WriteErrorJSON(w, http.StatusInternalServerError, "internal server error")
		return
	}

	log.Printf("Info: UpdateSubscription#ServeHTTP: updated subscription into %v\n", sub)

	httputil.WriteJSON(w, http.StatusOK, sub)
}
