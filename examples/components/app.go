package components

import (
	"github.com/factorapp/factor/examples/models"
)

type App struct {
	//factor.Router
	Todos []*models.Todo
	Page  string
}
