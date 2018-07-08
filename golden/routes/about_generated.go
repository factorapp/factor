// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package routes

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
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
	)
}
