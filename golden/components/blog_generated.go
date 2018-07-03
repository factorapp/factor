
package components

import (
	
	"github.com/bketelsen/factor/markup"
)



var BlogTemplate =  `<main>
    <Nav />
    <div>Brian Was Here</div>
</main>` 
var BlogStyles = `
    main {
        position: relative;
        max-width: 56em;
        background-color: white;
        padding: 2em;
        margin: 0 auto;
        box-sizing: border-box;
    }
`


func (t *Blog) Render() string {
	return BlogTemplate
}

func init() {
	markup.Register(&Blog{})
}
