package codegen

import (
	"bytes"
	"text/template"
)

// ClientGoMain returns the Go code for the main function for the wasm app
func ClientGoMain(appPath string) (string, error) {
	b := new(bytes.Buffer)
	data := map[string]string{"AppPath": appPath}
	if err := clMainTemplate.Execute(b, data); err != nil {
		return "", err
	}
	return string(b.Bytes()), nil
}

var clMainTemplate = template.Must(template.New("cl").Parse(clMainTemplateStr))

const clMainTemplateStr = `// build +js,wasm
package main

import (
	"{{.AppPath}}/routes"
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
