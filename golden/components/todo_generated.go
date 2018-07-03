package components

import (
	"html/template"
	"os"

	"github.com/bketelsen/factor/markup"
)

var TodoTemplate = `<main>
	<Nav />
	<div>
	    <h1>{{ .Todo.Get.Name }}</h1>
	    <small>{{ .Todo.Get.Description }}</small>
	    <div>(<a href="{{ .Todo.Permalink }}">Permalink</a>)</div>
	</div>
</main>`
var TodoStyles = ``

func (t *Todo) Render() string {
	tpl := template.Must(template.New("todo").Parse(TodoTemplate))
	return tpl.Execute(
		/*what do you put in here for the writer?*/
		os.Stdout,
		t,
	)
}

func init() {
	markup.Register(&Todo{})
}
