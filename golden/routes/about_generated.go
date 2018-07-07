// This file was created with https://jsgo.io/dave/html2vecty
package routes

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

type About struct {
	vecty.Core
}

func (p *About) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Heading1(
			vecty.Text("About this site"),
		),
		elem.Paragraph(
			vecty.Text("This is the 'about' page. There's not much here."),
		),

		elem.Anchor(
			vecty.Markup(
				prop.Href("/"),
			),
			vecty.Text("Permalink"),
		),
	)
}
