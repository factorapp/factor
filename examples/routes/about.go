package routes

import "github.com/gopherjs/vecty"
import "github.com/gopherjs/vecty/elem"

func asAnArg(arg *vecty.HTML) *vecty.HTML {
	return arg
}

type About struct{}

func (p *About) Render() vecty.ComponentOrHTML {

	header := elem.Heading1(vecty.Text("Hello"))
	content := elem.Paragraph(vecty.

		// hello world I'm a comment
		Text("paragraphs"))

	return elem.Div(vecty.Markup(vecty.Class("content")), vecty.Text(header), vecty.Text(content), vecty.Text(asAnArg(elem.Paragraph(vecty.Text("help I'm caught in the web")))))

}
