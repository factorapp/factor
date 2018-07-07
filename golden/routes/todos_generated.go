// This file was created with https://jsgo.io/dave/html2vecty
package main

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

type Todos struct {
	vecty.Core
}

func (p *Todos) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("todos"),
			),
			vecty.Text("{{ range .Todos }}"),
			vecty.Text("{{ .TodoComponent . }}\n    {{ end }}"),
		),
	)
}
