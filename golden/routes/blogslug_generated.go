
package components

import (
	
	"github.com/bketelsen/factor/markup"
)

type Blogslug struct {

}

var BlogslugTemplate =  `<h1>blog</h1>

<p>This is a blog page, will populate by parameter</p>` 
var BlogslugStyles = ``


func (t *Blogslug) Render() string {
	return BlogslugTemplate
}

func init() {
	markup.Register(&Blogslug{})
}
