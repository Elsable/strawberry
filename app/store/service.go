package store

import (
	"github.com/google/uuid"
)

// Service wraps logic on top of storage
type Service struct {
	Engine Engine
}

func (s *Service) Create(r Resource) (resourceID string, err error) {

	r.ID = uuid.New().String()
	return s.Engine.Create(r)
}

func (s *Service) Get(resourceID string) (r Resource, err error) {
	return s.Engine.Get(resourceID)
}

func (s *Service) List(limit int) (list *[]Resource, err error) {
	return s.Engine.List(limit)
}

func (s *Service) Delete(resourceID string) (err error) {
	return s.Delete(resourceID)
}
