package subscription

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	db *pgx.Conn
}

type ErrorSubscriptionNotFound struct {
	message string
}

func NewStorage(db *pgx.Conn) *Storage {
	return &Storage{db}
}

func (e *ErrorSubscriptionNotFound) Error() string {
	return e.message
}

// could return ErrorSubscriptionNotFound
// could return some other errors from pgx
func (s *Storage) GetSubscription(id int) (*Subscription, error) {
	sub := &Subscription{}
	err := s.db.QueryRow(
		context.Background(),
		`SELECT id,
		        service_name,
				price,
				user_id,
				start_date,
				end_date
		FROM subscriptions WHERE id = $1`,
		id,
	).Scan(
		&sub.ID,
		&sub.ServiceName,
		&sub.Price,
		&sub.UserID,
		&sub.StartDate,
		&sub.EndDate,
	)
	if err == pgx.ErrNoRows {
		return nil, &ErrorSubscriptionNotFound{fmt.Sprintf("GetSubscription: record with id %d not found", id)}
	}
	if err != nil {
		return nil, err
	}

	return sub, nil
}
