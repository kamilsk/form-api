package service

import "github.com/kamilsk/form-api/pkg/domain"

// Storage defines the behavior of Data Access Object.
type Storage interface {
	// AddData inserts form data and returns their ID.
	AddData(domain.UUID, map[string][]string) (int64, error)
	// Schema returns the form schema with provided UUID.
	Schema(domain.UUID) (domain.Schema, error)
	// UUID returns a new generated unique identifier.
	UUID() (domain.UUID, error)
}
