package component

var comptpl = `
package components

import (
	{{.WriteImports}}
	"github.com/bketelsen/factor/markup"
)

var {{.Name}}Template =  {{.QuotedTemplate}} 
var {{.Name}}Styles = {{.QuotedStyle}}

type {{.Name}} struct{

}

func (t *{{.Name}}) Render() {
	return {{.Name}}Template
}

func init() {
	markup.Register(&{{.Name}}{})
}
`
