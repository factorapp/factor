package transpiler

import (
	"encoding/xml"
	"strings"

	"github.com/dave/jennifer/jen"
)

func ComponentElement(appPackage, componentName string, token *xml.StartElement) *jen.Statement {
	var component string
	var qual bool
	vectyPackage := appPackage + "/components"
	// vectyFunction = component
	// vectyParamater = tag
	qual = true
	component = strings.TrimLeft(token.Name.Local, "components.")
	if qual {
		baseDecl := jen.Id("&").Qual(vectyPackage, "").Id(component).Values(jen.DictFunc(func(d jen.Dict) {
			for _, v := range token.Attr {
				d[jen.Id(v.Name.Local)] = jen.Lit(v.Value)
			}
		}))
		return baseDecl
	}
	baseDecl := jen.Id("&").Id(component).Values(jen.DictFunc(func(d jen.Dict) {
		for _, v := range token.Attr {
			d[jen.Id(v.Name.Local)] = jen.Lit(v.Value)
		}
	}))
	return baseDecl

}
