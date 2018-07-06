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
			vecty.Text("About this site"),
		),
		elem.Paragraph(
			vecty.Text("This is the 'about' page. There's not much here."),
		),
	)
}
