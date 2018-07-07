// This file was created with https://jsgo.io/dave/html2vecty
package main

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

type Nav struct {
	vecty.Core
}

func (p *Nav) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Navigation(
			elem.UnorderedList(
				elem.ListItem(
					vecty.Text("This is an item"),
				),
				elem.ListItem(
					vecty.Text("This is another item"),
				),
			),
		),
	)
}
