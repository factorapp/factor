// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package components

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

type Todo struct {
	vecty.Core
}

func (p *Todo) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			elem.Heading1(
				vecty.Text("vecty-field:Name"),
			),
			elem.Small(
				vecty.Text("vecty-field:Description"),
			),
			elem.Div(
				vecty.Text("("),
				elem.Anchor(
					vecty.Markup(
						prop.Href("{vecty-field:Permalink}"),
					),
					vecty.Text("Permalink"),
				),
				vecty.Text(")"),
			),
			elem.Div(

				vecty.Text("You are "),
				vecty.Text(p.GetAge()),
				vecty.Text(" years old"),
			),
			elem.Div(

				vecty.Text("gary"),
				vecty.Text(p.GetSomething()),
				vecty.Text("susan"),
				vecty.Text(p.GetSomething()),
				vecty.Text("bob"),
			),
		),
	)
}
