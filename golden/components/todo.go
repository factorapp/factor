package components

import "github.com/gowasm/vecty"

type Todo struct {
	vecty.Core
	Name        string
	Description string
	Permalink   string
}

func (t *Todo) GetAge() string {
	return "5"
}
