package models

import (
	"math/rand"
	"time"

	"github.com/satori/go.uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Todo is a model
type Todo struct {
	ID          uuid.UUID
	Name        string
	Description string
	Permalink   string
}

func (t Todo) GetAge() int {
	return rand.Int()
}
