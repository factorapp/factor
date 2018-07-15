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
		baseDecl := jen.Id("&").Add(jen.Qual(vectyPackage, "").Add(
			jen.Id(component),
		))
		attrCall := jen.Options{
			Close:     "",
			Multi:     false,
			Open:      "",
			Separator: ",",
		}
		block := jen.CustomFunc(attrCall, func(g *jen.Group) {
			for _, v := range token.Attr {
				g.Id(v.Name.Local).Id(":").Lit(v.Value).Id(",")
			}
		})

		baseDecl.Block(block)
		return baseDecl
	}
	baseDecl := jen.Id("&").Add(
		jen.Id(component),
	)
	attrCall := jen.Options{
		Close:     "",
		Multi:     true,
		Open:      "",
		Separator: ",",
	}
	block := jen.CustomFunc(attrCall, func(g *jen.Group) {
		for _, v := range token.Attr {
			g.Id(v.Name.Local).Id(":").Lit(v.Value).Id(",")
		}
	})
	baseDecl.Block(block)

	return baseDecl

}
