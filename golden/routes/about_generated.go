package routes

import (
	"github.com/bketelsen/factor/markup"
)

type About struct {
}

var AboutTemplate = `<h1>About this site</h1>

<p>This is the 'about' page. There's not much here.</p>`
var AboutStyles = ``

func (t *About) Render() string {
	return AboutTemplate
}

func init() {
	markup.Register(&About{})
}
