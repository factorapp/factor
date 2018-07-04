package routes

import (
	"github.com/bketelsen/factor/markup"
)

type About struct {
}

var AboutTemplate = `<main><h1>About this site</h1>

<p>This is the 'about' page. There's not much here.</p></main>`
var AboutStyles = ``

func (t *About) Render() string {
	return AboutTemplate
}
func (t *About) Style() string {
	return AboutStyles
}
func init() {
	markup.Register(&About{})
}
