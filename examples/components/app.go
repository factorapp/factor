package components

import (
	"github.com/factorapp/factor/golden/models"
)

type App struct {
	//factor.Router
	Todos []*models.Todo
	Page  string
}
