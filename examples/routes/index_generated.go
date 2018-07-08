// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package routes

import (
	components "github.com/factorapp/factor/examples/components"
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

func (p *Index) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Body(
			&components.Nav{},
		),
		elem.Main(
			vecty.Markup(
				vecty.Attribute("role", "main"),
				vecty.Class("container"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("starter-template"),
				),
				elem.Heading1(
					vecty.Text("Bootstrap starter template"),
				),
				elem.Paragraph(
					vecty.Markup(
						vecty.Class("lead"),
					),
					vecty.Text("Use this document as a way to quickly start any new project."),
					elem.Break(),
					vecty.Text("All you get is this text and a mostly barebones HTML document."),
				),
			),
		),
	)
}
