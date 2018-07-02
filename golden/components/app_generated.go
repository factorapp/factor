
package components

import (
	
	"github.com/bketelsen/factor/markup"
)



var AppTemplate =  `<main>
<Nav />
	<div>Brian Was Here</div>
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
	return AppTemplate
}

func init() {
	markup.Register(&App{})
}
