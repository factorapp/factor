// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package components

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

func (p *Todo) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			elem.Heading1(

				vecty.Text(p.Name),
			),
			elem.Small(

				vecty.Text(p.Description),
			),
			elem.Div(
				vecty.Text("("),
				elem.Anchor(
					vecty.Markup(
						prop.Href(p.Permalink),
					),
					vecty.Text("Permalink"),
				),
				vecty.Text(")"),
			),
			elem.Div(

				vecty.Text("Active for "),
				vecty.Text(p.GetAge()),
				vecty.Text(" days"),
			),
		),
	)
}
