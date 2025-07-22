package subscription

import (
	"github.com/lenkton/effective-mobile-test/pkg/httputil"
)

type Subscription struct {
	ID          int           `json:"id"`
	ServiceName string        `json:"service_name"`
	Price       int           `json:"price"`
	UserID      string        `json:"user_id"`
	StartDate   httputil.Time `json:"start_date"`
	EndDate     httputil.Time `json:"end_date,omitzero"`
}
