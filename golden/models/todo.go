package models

import (
	"github.com/satori/go.uuid"
)

type Todo struct {
	ID          uuid.UUID
	Name        string
	Description string
}

// AddTodo is the server-side implementation to add a Todo to the
// database
//
// NewTodo is a generated Todo. It is the same as Todo but also have a
// context.Context in it. Maybe this?
//
//	type NewTodo struct {
//		Ctx context.Context
//		ID uuid.UUID
//		Name string
//		Description  string
//	}
//
// TodoResp has some kind of error or status thing in it. Maybe this?
//
//	type TodoResp struct {
//		Result error
//	}
//
// note: I feel like this should be in a separate package that holds
// all the server stuff. I'll leave it here for now
func AddTodo(arg NewTodo, retVal TodoResp) error {
	return database.AddTodo(arg.Ctx)
}
