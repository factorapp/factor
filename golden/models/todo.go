package models

import (
	"github.com/satori/go.uuid"
)

// Todo is a model
type Todo struct {
	ID          uuid.UUID
	Name        string
	Description string
}

// AddTodo is the server-side implementation to add a Todo to the
// database.
//
// note: I feel like this should be in a separate package that holds
// all the server stuff. I'll leave it here for now
func AddTodo(arg NewTodo, retVal TodoResp) error {
	// database is something that hooks up to a persistent database
	return database.AddTodo(arg.Ctx)
}
