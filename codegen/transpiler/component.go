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
	// I'm not sure what qual was intended to mean (it's always true now) but it looks like perhaps you're
	// trying to avoid using Qual if the package path == local path? If so, no need! Qual handles this
	// gracefully... See: https://github.com/dave/jennifer#qual
	if qual {
		baseDecl := jen.Op("&").Qual(vectyPackage, component).Values(jen.DictFunc(func(d jen.Dict) {
			for _, v := range token.Attr {
				d[jen.Id(v.Name.Local)] = jen.Lit(v.Value)
			}
		}))
		return baseDecl
	}
	baseDecl := jen.Op("&").Id(component).Values(jen.DictFunc(func(d jen.Dict) {
		for _, v := range token.Attr {
			d[jen.Id(v.Name.Local)] = jen.Lit(v.Value)
		}
	}))
	return baseDecl

}
