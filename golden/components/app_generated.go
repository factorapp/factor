// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package components

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

func (p *App) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Main(
			&Nav{},
		),
		elem.Div(
			vecty.Text("Brian Was Here"),
		),
	)
}
