package component

var comptpl = `
package components

import (
	{{.WriteImports}}
	"github.com/factorapp/factor/markup"
)

{{ if .Struct }}type {{.Name}} struct {

}{{end}}

var {{.Name}}Template =  {{.QuotedTemplate}} 
var {{.Name}}Styles = {{.QuotedStyle}}


func (t *{{.Name}}) Render() string {
	return {{.Name}}Template
}

func init() {
	markup.Register(&{{.Name}}{})
}
`
