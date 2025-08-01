package subscription

import (
	"context"
	"fmt"
	"log"

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

// could return some errors from pgx
func (s Storage) ListSubscriptions() ([]*Subscription, error) {
	rows, err := s.db.Query(context.Background(), "select * from subscriptions")
	if err != nil {
		log.Printf("Error: ListSubscriptions#ServerHTTP/Query: %v\n", err)
		return []*Subscription{}, fmt.Errorf("ListSubscriptions: Query: %v", err)
	}

	subs, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[Subscription])
	if err != nil {
		return []*Subscription{}, fmt.Errorf("ListSubscriptions: CollectRows: %v", err)
	}

	return subs, nil
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

// could return some other errors from pgx
func (s *Storage) CreateSubscription(sub *Subscription) (int, error) {
	var id int
	err := s.db.QueryRow(
		context.Background(),
		`INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		sub.ServiceName,
		sub.Price,
		sub.UserID,
		sub.StartDate.Time,
		sub.EndDate,
	).Scan(&id)

	return id, err
}
