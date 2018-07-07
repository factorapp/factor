// This file was created with https://jsgo.io/dave/html2vecty
package components

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

/*
type App struct {
	vecty.Core
}
*/

func (p *App) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Main(
			vecty.Tag(
				"Nav",
			),
			elem.Div(
				vecty.Text("Brian Was Here"),
			),
		),
	)
}
