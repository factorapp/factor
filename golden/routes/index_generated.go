// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package routes

import (
	components "github.com/factorapp/factor/golden/components"
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/event"
	"github.com/gowasm/vecty/prop"
)

func (p *Index) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Main(
			&components.Nav{

				MyProp: "blue",
			},
		),
		elem.Heading1(
			vecty.Text("Great success!"),
		),
		elem.Figure(
			elem.Image(
				vecty.Markup(
					vecty.Attribute("alt", "Borat"),
					prop.Src("/assets/great-success.png"),
				),
			),
			elem.FigureCaption(
				vecty.Text("HIGH FIVE!"),
			),
		),
		elem.Paragraph(
			elem.Input(
				vecty.Markup(
					prop.Type(prop.TypeSubmit),
					prop.Value("increment counter"),
					event.Click(p.OnClick),
				),
			),
			elem.Strong(

				vecty.Text(p.CountText),
			),
		),
	)
}
