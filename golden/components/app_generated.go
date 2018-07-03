package components

import (
	"bytes"
	"html/template"
	"log"

	"github.com/bketelsen/factor/markup"
)

var AppTemplate = `<main>
    <div>{{.Greeting}}</div>
	<a onclick="OnClick" href="/blog">Blog</a>
</main>`
var AppStyles = `
    main {
        position: relative;
        max-width: 56em;
        background-color: white;
        padding: 2em;
        margin: 0 auto;
        box-sizing: border-box;
    }
`

func (t *App) Render() string {
	b := new(bytes.Buffer)
	tpl, err := template.New("app").Parse(AppTemplate)
	if err != nil || tpl == nil {
		log.Printf("couldn't parse app template")
		return ""
	}
	log.Printf("parsed app template, continuing")
	tpl.Execute(b, t)
	return b.String()
}

func init() {
	markup.Register(&App{})
}
