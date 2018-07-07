// This file was created with https://jsgo.io/dave/html2vecty
package routes

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

type Blogslug struct {
	vecty.Core
}

func (p *Blogslug) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Heading1(
			vecty.Text("blog"),
		),
		elem.Paragraph(
			vecty.Text("This is a blog page, will populate by parameter"),
		),
	)
}
