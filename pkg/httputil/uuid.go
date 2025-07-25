package httputil

import "github.com/google/uuid"

func NewUUID() UUID {
	return UUID{uuid.New()}
}

type UUID struct {
	uuid.UUID
}
