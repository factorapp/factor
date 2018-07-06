// This file was created with https://jsgo.io/dave/html2vecty
package main

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

func main() {
	vecty.RenderBody(&Page{})
}

type Page struct {
	vecty.Core
}

func (p *Page) Render() vecty.ComponentOrHTML {
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
		),
	)
}
