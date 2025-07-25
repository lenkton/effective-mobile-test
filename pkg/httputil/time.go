package httputil

import (
	"database/sql/driver"
	"fmt"
	"log"
	"time"
)

type Time struct {
	time.Time
}

// could i use a pointer??
func (t Time) Value() (driver.Value, error) {
	return t.Time, nil
}

func (t *Time) Scan(src any) error {
	switch src := src.(type) {
	case time.Time:
		t.Time = src
		return nil
	default:
		err := fmt.Errorf("error: cannot handle type: %T", src)
		// TODO: prettify log message
		log.Println(err)
		return err
	}
}

func (t *Time) UnmarshalJSON(src []byte) error {
	parsedTime, err := time.Parse("01-2006", string(src))
	t.Time = parsedTime
	return err
}

func (t Time) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, "\"%s\"", t.Format("01-2006")), nil
}
