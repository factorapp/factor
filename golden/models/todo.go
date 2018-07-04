package models

import (
	"github.com/satori/go.uuid"
)

// Todo is a model
type Todo struct {
	ID          uuid.UUID
	Name        string
	Description string
	Permalink   string
}
