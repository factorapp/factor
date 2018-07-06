// build +js,wasm
package main

import (
	"github.com/bketelsen/factor/golden/components"
	"github.com/gowasm/vecty"
)

func main() {
	c := make(chan struct{}, 0)

	vecty.SetTitle("Markdown Demo")
	vecty.RenderBody(&components.PageView{
		Input: `# Markdown Example
	
	This is a live editor, try editing the Markdown on the right of the page.
	`,
	})
	<-c
}
