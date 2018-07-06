// This file was created with https://jsgo.io/dave/html2vecty
package main

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

func main() {
	vecty.RenderBody(&Page{})
}

type Page struct {
	vecty.Core
}

func (p *Page) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Heading1(
			vecty.Text("Todo"),
		),
		elem.Paragraph(
			vecty.Text("I'm not sure exactly how this'll work yet. Haven't run factor dev against it yet..."),
		),
	)
}
