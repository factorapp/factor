package cmd

var clMainTemplate = `// build +js,wasm
package main

import (
	"github.com/factorapp/factor/examples/routes"
	"github.com/gowasm/router"
	"github.com/gowasm/vecty"
)

func main() {
	c := make(chan struct{}, 0)
	// Create a new Router object
	r := router.New()
	//r.ShouldInterceptLinks = true
	// Use HandleFunc to add routes.
	r.HandleFunc("/", func(context *router.Context) {

		// The handler for this route simply grabs the name parameter
		// from the map of params and says hello.
		vecty.SetTitle("Factor: Home")
		vecty.RenderBody(&routes.Index{})
	})
	r.HandleFunc("/about", func(context *router.Context) {

		// The handler for this route simply grabs the name parameter
		// from the map of params and says hello.
		vecty.SetTitle("Factor: About")
		vecty.RenderBody(&routes.About{})
	})
	// You must call Start in order to start listening for changes
	// in the url and trigger the appropriate handler function.
	r.Start()
	<-c
}
`
