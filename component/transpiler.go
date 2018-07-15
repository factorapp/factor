package component

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"regexp"

	"errors"

	"bytes"
	"io"

	"strings"

	"encoding/xml"

	"github.com/aymerick/douceur/parser"
	"github.com/dave/jennifer/jen"
	"github.com/factorapp/factor/codegen/transpiler"
)

var callRegexp = regexp.MustCompile(`{vecty-call:([a-zA-Z0-9_\-]+)}`)
var fieldRegexp = regexp.MustCompile(`{vecty-field:([a-zA-Z0-9_\-]+})`)

func NewTranspiler(r io.ReadCloser, createStruct bool, appPackage, componentName, packageName string) (*Transpiler, error) {
	s := &Transpiler{
		reader:        r,
		createStruct:  createStruct,
		appPackage:    appPackage,
		packageName:   packageName,
		componentName: componentName,
	}
	err := s.read()
	if err != nil {
		return s, err
	}
	err = s.transcode()
	if err != nil {
		return s, err
	}
	return s, nil
}

type Transpiler struct {
	reader        io.ReadCloser
	createStruct  bool
	appPackage    string
	componentName string
	packageName   string
	html, code    string
}

func (s *Transpiler) read() error {
	bb, err := ioutil.ReadAll(s.reader)
	if err != nil {
		return errors.New("reading component template")
	}
	s.html = string(bb)
	return nil
}

func (s *Transpiler) Code() string {
	return s.code
}

func (s *Transpiler) transcode() error {
	// check for valid HTML
	if _, err := template.New("syntaxcheck").Parse(s.html); err != nil {
		return err
	}

	decoder := xml.NewDecoder(bytes.NewBufferString(s.html))

	EOT := errors.New("end of tag")
	call := jen.Options{
		Close:     ")",
		Multi:     true,
		Open:      "(",
		Separator: ",",
	}
	/*values := jen.Options{
		Close:     "}",
		Multi:     true,
		Open:      "{",
		Separator: ",",
	}*/

	var transcode func(*xml.Decoder) ([]jen.Code, error)
	transcode = func(decoder *xml.Decoder) (code []jen.Code, err error) {
		token, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch token := token.(type) {
		case xml.StartElement:
			tag := token.Name.Local
			vectyFunction, ok := transpiler.ElemNames[tag]
			vectyPackage := "github.com/gowasm/vecty/elem"
			vectyParamater := ""
			var ce *jen.Statement
			if !ok {
				if strings.HasPrefix(token.Name.Space, "components") {
					// not sure if we need this?
					componentName := strings.TrimLeft(tag, "components.")
					ce = transpiler.ComponentElement(s.appPackage, componentName, &token)
					vectyFunction = ""
					//ce.Render(os.Stdout)
					//return ce, nil
					//jen.Add(ce)
				} else {
					vectyFunction = "Tag"
					vectyPackage = "github.com/gowasm/vecty"
					vectyParamater = tag
				}
			}
			var outer error

			q := jen.Qual(vectyPackage, vectyFunction).CustomFunc(call, func(g *jen.Group) {
				if vectyParamater != "" {
					g.Lit(vectyParamater)
				}
				if ce == nil && len(token.Attr) > 0 {
					g.Qual("github.com/gowasm/vecty", "Markup").CustomFunc(call, func(g *jen.Group) {
						for _, v := range token.Attr {
							switch {
							case v.Name.Local == "style":
								css, err := parser.ParseDeclarations(v.Value)
								if err != nil {
									outer = err
									return
								}
								for _, dec := range css {
									if dec.Important {
										dec.Value += "!important"
									}
									g.Qual("github.com/gowasm/vecty", "Style").Call(
										jen.Lit(dec.Property),
										jen.Lit(dec.Value),
									)
								}
							case v.Name.Local == "class":
								g.Qual("github.com/gowasm/vecty", "Class").CallFunc(func(g *jen.Group) {
									classes := strings.Split(v.Value, " ")
									for _, class := range classes {
										g.Lit(class)
									}
								})
							case strings.HasPrefix(v.Name.Local, "data-"):
								attribute := strings.TrimPrefix(v.Name.Local, "data-")
								g.Qual("github.com/gowasm/vecty", "Data").Call(
									jen.Lit(attribute),
									jen.Lit(v.Value),
								)

							case transpiler.BoolProps[v.Name.Local] != "":
								value := v.Value == "true"
								g.Qual("github.com/gowasm/vecty/prop", transpiler.BoolProps[v.Name.Local]).Call(
									jen.Lit(value),
								)
							case transpiler.StringProps[v.Name.Local] != "":
								if strings.HasPrefix(v.Value, "{vecty-field:") {
									field := strings.TrimLeft(v.Value, "{vecty-field:")
									field = field[:len(field)-1]
									g.Qual("github.com/gowasm/vecty/prop", transpiler.StringProps[v.Name.Local]).Call(
										jen.Id("p." + field),
									)
								} else {
									g.Qual("github.com/gowasm/vecty/prop", transpiler.StringProps[v.Name.Local]).Call(
										jen.Lit(v.Value),
									)
								}
							case strings.HasPrefix(v.Name.Space, "vecty"):
								field := strings.TrimLeft(v.Name.Local, "on")
								field = strings.ToLower(field)
								g.Qual("github.com/gowasm/vecty/event", strings.Title(field)).Call(
									jen.Id("p." + v.Value),
								)
							case strings.HasPrefix(v.Name.Space, "components"):
								component := strings.TrimLeft(v.Name.Local, "components.")
								jen.Id("&" + component + "{}")
							case v.Name.Local == "xmlns":
								g.Qual("github.com/gowasm/vecty", "Namespace").Call(
									jen.Lit(v.Value),
								)
							case v.Name.Local == "type" && transpiler.TypeProps[v.Value] != "":
								g.Qual("github.com/gowasm/vecty/prop", "Type").Call(
									jen.Qual("github.com/gowasm/vecty/prop", transpiler.TypeProps[v.Value]),
								)
							default:
								g.Qual("github.com/gowasm/vecty", "Attribute").Call(
									jen.Lit(v.Name.Local),
									jen.Lit(v.Value),
								)
							}
						}
					})
				}
				for {
					c, err := transcode(decoder)
					if err != nil {
						if err == EOT {
							break
						}
						outer = err
						return
					}
					if c != nil {
						g.Add(c...)
					}
				}
			})
			if outer != nil {
				return nil, outer
			}
			if ce != nil {
				return []jen.Code{ce}, nil
			}
			return []jen.Code{q}, nil
		case xml.CharData:
			str := string(token)
			hasCall := callRegexp.MatchString(str)
			hasField := fieldRegexp.MatchString(str)
			hasSpecial := hasCall || hasField

			if hasSpecial {
				/*
					fieldQualifier := func(name string) (*jen.Statement, error) {

						fmt.Println("field qualifier:", str, name)
						n := strings.TrimLeft(name, "{vecty-field:")
						n = strings.TrimRight(n, "}")
						return jen.Qual("github.com/gowasm/vecty", "Text").Call(
							// TODO: struct qualifier
							jen.Id("p." + n),
						), nil
					}
				*/
				/*
					callQualifier := func(lhs, name, rhs string) (*jen.Statement, error) {
						fmt.Println("call qualifier:", str, name)
						fnCall := strings.TrimLeft(name, "{vecty-call:")
						fnCall = strings.Replace(fnCall, "}", "", -1)
						stmt := jen.Qual("github.com/gowasm/vecty", "Text").Call(
							jen.Lit(lhs),
						)
						stmt.Qual("github.com/gowasm/vecty", "Text").Call(
							jen.Id("p." + fnCall + "()"),
						)
						stmt.Qual("github.com/gowasm/vecty", "Text").Call(
							jen.Lit(rhs),
						)
						return stmt, nil

					}
				*/
				// TODO: find next index for each field and call regexp.
				// build up a string of text statements for static text, and
				// generated calls or field accesses
				if hasCall {

					/*
						re := regexp.MustCompile(`({vecty-call:[a-zA-Z0-9]+})`)
						t := "{vecty-call:Alone}thing{vecty-call:Thing}stuff{vecty-call:Stuff}other{vecty-call:Blue}"
						crResult := re.FindAllStringIndex(t, -1)
						index := 0
						for matchNumber, match := range crResult {
							fmt.Println(match)
							fmt.Println("Before:", t[index:match[0]])
							fmt.Println("Match:", t[match[0]:match[1]])
							fmt.Println(re.FindAllStringSubmatch(t, index))
							if matchNumber < len(crResult)-1 {
								// there's another match

								fmt.Println(matchNumber, len(crResult), index, crResult, "There is another")
								fmt.Println("next Result:",crResult[matchNumber+1][0])
								fmt.Println(t[match[1]:])
								fmt.Println(crResult[matchNumber+1])
								fmt.Println("Internal After:", t[match[1]:crResult[matchNumber+1][0]])
							}

							fmt.Println("After:", t[match[1]:])
							index = match[1]
						}
					*/
					var statements []jen.Code
					crResult := callRegexp.FindAllStringIndex(str, -1)
					index := 0
					for matchNumber, match := range crResult {
						var before, between, after string
						before = str[index:match[0]]
						fnCall := str[match[0]:match[1]]
						fnCall = strings.TrimLeft(fnCall, "{vecty-call:")
						fnCall = strings.Replace(fnCall, "}", "", -1)
						if matchNumber < len(crResult)-1 {
							// there's another match
							between = str[match[1]:crResult[matchNumber+1][0]]
						}
						after = str[match[1]:]
						/*
							g.Qual("fmt", "Sprintf").Call(
								jen.Lit("%s%s%s"),
								jen.Lit(lhs),
								jen.Id("p."+fnCall+"()"),
								jen.Lit(rhs),
							)
						*/
						if before != "" && !strings.Contains(before, "vecty-call") {
							statements = append(statements, jen.Qual("github.com/gowasm/vecty", "Text").Call(
								jen.Lit(before),
							))
						}
						statements = append(statements, jen.Qual("github.com/gowasm/vecty", "Text").Call(
							jen.Id("p."+fnCall+"()"),
						))
						if between != "" && !strings.Contains(between, "vecty-call") {
							statements = append(statements, jen.Qual("github.com/gowasm/vecty", "Text").Call(
								jen.Lit(between),
							))
						}
						if after != "" && !strings.Contains(after, "vecty-call") {
							statements = append(statements, jen.Qual("github.com/gowasm/vecty", "Text").Call(
								jen.Lit(after),
							))
						}
					}
					return statements, nil

				}
				if hasField {

					/*
						re := regexp.MustCompile(`({vecty-call:[a-zA-Z0-9]+})`)
						t := "{vecty-call:Alone}thing{vecty-call:Thing}stuff{vecty-call:Stuff}other{vecty-call:Blue}"
						crResult := re.FindAllStringIndex(t, -1)
						index := 0
						for matchNumber, match := range crResult {
							fmt.Println(match)
							fmt.Println("Before:", t[index:match[0]])
							fmt.Println("Match:", t[match[0]:match[1]])
							fmt.Println(re.FindAllStringSubmatch(t, index))
							if matchNumber < len(crResult)-1 {
								// there's another match

								fmt.Println(matchNumber, len(crResult), index, crResult, "There is another")
								fmt.Println("next Result:",crResult[matchNumber+1][0])
								fmt.Println(t[match[1]:])
								fmt.Println(crResult[matchNumber+1])
								fmt.Println("Internal After:", t[match[1]:crResult[matchNumber+1][0]])
							}

							fmt.Println("After:", t[match[1]:])
							index = match[1]
						}
					*/
					var statements []jen.Code
					crResult := fieldRegexp.FindAllStringIndex(str, -1)
					index := 0
					for matchNumber, match := range crResult {
						var before, between, after string
						before = str[index:match[0]]
						field := str[match[0]:match[1]]
						field = strings.TrimLeft(field, "{vecty-field:")
						field = strings.Replace(field, "}", "", -1)
						if matchNumber < len(crResult)-1 {
							// there's another match
							between = str[match[1]:crResult[matchNumber+1][0]]
						}
						after = str[match[1]:]
						/*
							g.Qual("fmt", "Sprintf").Call(
								jen.Lit("%s%s%s"),
								jen.Lit(lhs),
								jen.Id("p."+fnCall+"()"),
								jen.Lit(rhs),
							)
						*/
						if before != "" && !strings.Contains(before, "vecty-field") {
							statements = append(statements, jen.Qual("github.com/gowasm/vecty", "Text").Call(
								jen.Lit(before),
							))
						}
						statements = append(statements, jen.Qual("github.com/gowasm/vecty", "Text").Call(
							jen.Id("p."+field),
						))
						if between != "" && !strings.Contains(between, "vecty-field") {
							statements = append(statements, jen.Qual("github.com/gowasm/vecty", "Text").Call(
								jen.Lit(between),
							))
						}
						if after != "" && !strings.Contains(after, "vecty-field") {
							statements = append(statements, jen.Qual("github.com/gowasm/vecty", "Text").Call(
								jen.Lit(after),
							))
						}
					}
					return statements, nil

				}

			}
			s := strings.TrimSpace(string(token))
			if s == "" {
				return nil, nil
			}
			return []jen.Code{jen.Qual("github.com/gowasm/vecty", "Text").Call(jen.Lit(s))}, nil
		case xml.EndElement:
			return nil, EOT
		case xml.Comment:
			return nil, nil
		default:
			fmt.Printf("%T %#v \n", token, token)
		}
		return nil, nil
	}

	file := jen.NewFile(s.packageName)
	file.PackageComment("This file was created with https://github.com/factorapp/factor")
	file.PackageComment("using https://jsgo.io/dave/html2vecty")
	file.ImportNames(map[string]string{
		"github.com/gowasm/vecty":                           "vecty",
		"github.com/gowasm/vecty/elem":                      "elem",
		"github.com/gowasm/vecty/prop":                      "prop",
		"github.com/gowasm/vecty/event":                     "event",
		"github.com/gowasm/vecty/style":                     "style",
		"_ github.com/factorapp/factor/examples/components": "components",
	})
	var elements []jen.Code
	for {
		c, err := transcode(decoder)
		if err != nil {
			if err == io.EOF || err == EOT {
				break
			}
			s.code = fmt.Sprintf("%s", err)
			return nil
		}
		if c != nil {
			elements = append(elements, c...)
		}
	}
	/*
		func main() {
			vecty.RenderBody(&Page{})
		}

		type Page struct {
			vecty.Core
		}

		func (*Page) Render() vecty.ComponentOrHTML {
			return elem.Body(...)
		}
	*/
	/*	file.Func().Id("main").Params().Block(
			jen.Qual("github.com/gowasm/vecty", "RenderBody").Call(
				jen.Op("&").Id("Page").Values(),
			),
		)
	*/
	if s.createStruct {
		file.Type().Id(s.componentName).Struct(
			jen.Qual("github.com/gowasm/vecty", "Core"),
		)
	}
	if s.packageName == "routes" || s.packageName == "pages" {
		file.Func().Params(jen.Id("p").Op("*").Id(s.componentName)).Id("Render").Params().Qual("github.com/gowasm/vecty", "ComponentOrHTML").Block(
			jen.Qual("github.com/gowasm/vecty", "SetTitle").Call(
				jen.Id("p.GetTitle()"),
			),
			jen.Return(
				// TODO: wrap in if - only body for a "route"
				jen.Qual("github.com/gowasm/vecty/elem", "Body").Custom(call, elements...),
			),
		)
	} else {
		file.Func().Params(jen.Id("p").Op("*").Id(s.componentName)).Id("Render").Params().Qual("github.com/gowasm/vecty", "ComponentOrHTML").Block(
			// TODO: wrap in if - only body for a "route"
			jen.Return(elements...),
		)
	}
	/*if len(elements) == 1 {
		file.Var().Id("Element").Op("=").Add(elements[0])
	} else if len(elements) > 1 {
		file.Var().Id("Elements").Op("=").Index().Op("*").Qual("github.com/gopherjs/vecty", "HTML").Custom(values, elements...)
	}*/

	buf := &bytes.Buffer{}
	if err := file.Render(buf); err != nil {
		s.code = fmt.Sprintf("%s", err)
		return nil
	}

	s.code = buf.String()
	return nil
}
