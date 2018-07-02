
package components

import (
	
	"github.com/bketelsen/factor/markup"
)

type Error struct {

}

var ErrorTemplate =  `<h1>status</h1>

<p>error.message</p>` 
var ErrorStyles = `
	h1,
	p {
		margin: 0 auto;
	}

	h1 {
		font-size: 2.8em;
		font-weight: 700;
		margin: 0 0 0.5em 0;
	}

	p {
		margin: 1em auto;
	}

	@media (min-width: 480px) {
		h1 {
			font-size: 4em;
		}
	}
`


func (t *Error) Render() string {
	return ErrorTemplate
}

func init() {
	markup.Register(&Error{})
}
