
package components

import (
	
	"github.com/bketelsen/factor/markup"
)

type Todo struct {

}

var TodoTemplate =  `<!-- maybe factor can somehow "bind" the models.Todo to this html? -->
<div>
    <h1>{{ .Todo.Get.Name }}</h1>
    <small>{{ .Todo.Get.Description }}</small>
    <div>(<a href="{{ .Todo.Permalink }}">Permalink</a>)</div>
</div>` 
var TodoStyles = ``


func (t *Todo) Render() string {
	return TodoTemplate
}

func init() {
	markup.Register(&Todo{})
}
