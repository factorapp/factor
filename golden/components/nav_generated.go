// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package components

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

func (p *Nav) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Navigation(
			elem.UnorderedList(
				elem.ListItem(

					vecty.Text("This a component property: "),
					vecty.Text(p.MyProp),
				),
				elem.ListItem(
					vecty.Text("This is another item"),
				),
			),
		),
	)
}
