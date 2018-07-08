// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package routes

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

type Todoid struct {
	vecty.Core
}

func (p *Todoid) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Heading1(
			vecty.Text("Todo"),
		),
		elem.Paragraph(
			vecty.Text("I'm not sure exactly how this'll work yet. Haven't run factor dev against it yet..."),
		),
	)
}
