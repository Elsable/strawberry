package store

import (
	"fmt"
	"time"
)

// Resource data
type Resource struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Location    Location  `json:"location"`
	Timestamp   time.Time `json:"time"`
}

func (r *Resource) String() string {
	return fmt.Sprintf("Resource ID=%s Type=%s Description=%s Location=%v Timestamp=%s", r.ID, r.Type, r.Description, r.Location, r.Timestamp)
}

// Location coordinates referenced from the GeoJSON
type Location struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

// Engine defines interface to save, load, remove and inc errors count for resources
type Engine interface {
	Create(r Resource) (resourceID string, err error)
	Get(resourceID string) (r Resource, err error)
	List(limit int) (list *[]Resource, err error)
	Delete(resourceID string) (err error)
}
