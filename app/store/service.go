package store

import (
	"github.com/google/uuid"
)

// Service wraps logic on top of storage
type Service struct {
	InMemory
}

func (s *Service) Create(r *Resource) (resourceID string, err error) {

	r.ID = uuid.New().String()
	return s.InMemory.Create(r)
}
