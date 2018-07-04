package routes

import (
	"github.com/bketelsen/factor/markup"
)

type Todoid struct {
}

var TodoidTemplate = `<h1>Todo</h1>

<p>I'm not sure exactly how this'll work yet. Haven't run factor dev against it yet...</p>`
var TodoidStyles = ``

func (t *Todoid) Render() string {
	return TodoidTemplate
}
func (t *Todoid) Style() string {
	return TodoidStyles
}
func init() {
	markup.Register(&Todoid{})
}
