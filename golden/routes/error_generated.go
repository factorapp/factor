// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package routes

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

type Error struct {
	vecty.Core
}

func (p *Error) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Heading1(
			vecty.Text("status"),
		),
		elem.Paragraph(
			vecty.Text("error.message"),
		),
		elem.Style(
			vecty.Text("h1,\n\tp {\n\t\tmargin: 0 auto;\n\t}\n\n\th1 {\n\t\tfont-size: 2.8em;\n\t\tfont-weight: 700;\n\t\tmargin: 0 0 0.5em 0;\n\t}\n\n\tp {\n\t\tmargin: 1em auto;\n\t}\n\n\t@media (min-width: 480px) {\n\t\th1 {\n\t\t\tfont-size: 4em;\n\t\t}\n\t}"),
		),
	)
}
