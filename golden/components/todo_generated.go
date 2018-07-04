package components

import (
	"github.com/bketelsen/factor/markup"
)

type Todo struct {
	Name        string
	Description string
	Permalink   string
}

var TodoTemplate = `<!-- maybe factor can somehow "bind" the models.Todo to this html? -->
<div>
    <h1>{{ .Name }}</h1>
    <small>{{ .Description }}</small>
    <div>(<a href="{{ .Permalink }}">Permalink</a>)</div>
</div>`
var TodoStyles = ``

func (t *Todo) Render() string {
	return TodoTemplate
}

func init() {
	markup.Register(&Todo{})
}
