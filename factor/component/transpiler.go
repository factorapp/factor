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
)

func NewTranspiler(r io.ReadCloser, createStruct bool, componentName, packageName string) (*Transpiler, error) {
	s := &Transpiler{
		reader:        r,
		createStruct:  createStruct,
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
func (s *Transpiler) Html() string {
	return s.html
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

	callRegexp := regexp.MustCompile(`{vecty-call:([a-zA-Z0-9_\-]+)}`)
	fieldRegexp := regexp.MustCompile(`{vecty-field:([a-zA-Z0-9_\-]+})`)

	var transcode func(*xml.Decoder) (jen.Code, error)
	transcode = func(decoder *xml.Decoder) (code jen.Code, err error) {
		token, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch token := token.(type) {
		case xml.StartElement:
			tag := token.Name.Local
			vectyFunction, ok := elemNameMap[tag]
			vectyPackage := "github.com/gowasm/vecty/elem"
			vectyParamater := ""
			if !ok {
				vectyFunction = "Tag"
				vectyPackage = "github.com/gowasm/vecty"
				vectyParamater = tag
			}
			var outer error
			q := jen.Qual(vectyPackage, vectyFunction).CustomFunc(call, func(g *jen.Group) {
				if vectyParamater != "" {
					g.Lit(vectyParamater)
				}
				if len(token.Attr) > 0 {
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

							case boolProps[v.Name.Local] != "":
								value := v.Value == "true"
								g.Qual("github.com/gowasm/vecty/prop", boolProps[v.Name.Local]).Call(
									jen.Lit(value),
								)
							case stringProps[v.Name.Local] != "":
								if strings.HasPrefix(v.Value, "{vecty-field:") {
									field := strings.TrimLeft(v.Value, "{vecty-field:")
									field = field[:len(field)-1]
									g.Qual("github.com/gowasm/vecty/prop", stringProps[v.Name.Local]).Call(
										jen.Id("p." + field),
									)
								} else {
									g.Qual("github.com/gowasm/vecty/prop", stringProps[v.Name.Local]).Call(
										jen.Lit(v.Value),
									)
								}
							case strings.HasPrefix(v.Name.Space, "vecty"):
								fmt.Println(v.Name.Space, v.Name.Local, v.Value)
								field := strings.TrimLeft(v.Name.Local, "on")
								field = strings.ToLower(field)
								g.Qual("github.com/gowasm/vecty/event", strings.Title(field)).Call(
									jen.Id("p." + v.Value),
								)

							case v.Name.Local == "xmlns":
								g.Qual("github.com/gowasm/vecty", "Namespace").Call(
									jen.Lit(v.Value),
								)
							case v.Name.Local == "type" && typeProps[v.Value] != "":
								g.Qual("github.com/gowasm/vecty/prop", "Type").Call(
									jen.Qual("github.com/gowasm/vecty/prop", typeProps[v.Value]),
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
						g.Add(c)
					}
				}
			})
			if outer != nil {
				return nil, outer
			}
			return q, nil
		case xml.CharData:
			str := string(token)
			hasCall := callRegexp.MatchString(str)
			hasField := fieldRegexp.MatchString(str)
			hasSpecial := hasCall || hasField

			if hasSpecial {
				fieldQualifier := func(name string) (*jen.Statement, error) {

					fmt.Println("field qualifier:", str, name)
					n := strings.TrimLeft(name, "{vecty-field:")
					n = strings.TrimRight(n, "}")
					return jen.Qual("github.com/gowasm/vecty", "Text").Call(
						// TODO: struct qualifier
						jen.Id("p." + n),
					), nil
				}
				/*
					callQualifier := func(lhs, name, rhs string) (*jen.Statement, error) {
						fmt.Println("call qualifier:", str, name)
						fnCall := strings.TrimLeft(name, "{vecty-call:")
						fnCall = strings.Replace(fnCall, "}", "", -1)
						stmt := jen.Add(
							jen.Qual("github.com/gowasm/vecty", "Text").Call(
								jen.Lit(lhs),
							))
						stmt.Add(jen.Qual("github.com/gowasm/vecty", "Text").Call(
							jen.Id("p." + fnCall + "()"),
						))
						stmt.Add(jen.Qual("github.com/gowasm/vecty", "Text").Call(
							jen.Lit(rhs),
						))
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
					callOpts := jen.Options{
						Close:     "",
						Multi:     true,
						Open:      "",
						Separator: ",",
					}
					q := jen.CustomFunc(callOpts, func(g *jen.Group) {
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
								g.Add(
									jen.Qual("github.com/gowasm/vecty", "Text").Call(
										jen.Lit(before),
									))
							}
							g.Add(
								jen.Qual("github.com/gowasm/vecty", "Text").Call(
									jen.Id("p." + fnCall + "()"),
								))
							if between != "" && !strings.Contains(between, "vecty-call") {
								g.Add(
									jen.Qual("github.com/gowasm/vecty", "Text").Call(
										jen.Lit(between),
									))
							}
							if after != "" && !strings.Contains(after, "vecty-call") {
								g.Add(
									jen.Qual("github.com/gowasm/vecty", "Text").Call(
										jen.Lit(after),
									))
							}
						}
					})
					return q, nil

				}
				if hasField {
					fmt.Println("Found a field")
					fqResult := fieldRegexp.FindString(str)
					fmt.Println("field result:", fqResult)
					return fieldQualifier((fqResult))
				}
			}
			s := strings.TrimSpace(string(token))
			if s == "" {
				return nil, nil
			}
			return jen.Qual("github.com/gowasm/vecty", "Text").Call(jen.Lit(s)), nil
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
		"github.com/gowasm/vecty":       "vecty",
		"github.com/gowasm/vecty/elem":  "elem",
		"github.com/gowasm/vecty/prop":  "prop",
		"github.com/gowasm/vecty/event": "event",
		"github.com/gowasm/vecty/style": "style",
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
			elements = append(elements, c)
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
	file.Func().Params(jen.Id("p").Op("*").Id(s.componentName)).Id("Render").Params().Qual("github.com/gowasm/vecty", "ComponentOrHTML").Block(
		jen.Return(
			jen.Qual("github.com/gowasm/vecty/elem", "Body").Custom(call, elements...),
		),
	)
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

var elemNameMap = map[string]string{
	"a":          "Anchor",
	"abbr":       "Abbreviation",
	"address":    "Address",
	"area":       "Area",
	"article":    "Article",
	"aside":      "Aside",
	"audio":      "Audio",
	"b":          "Bold",
	"base":       "Base",
	"bdi":        "BidirectionalIsolation",
	"bdo":        "BidirectionalOverride",
	"blockquote": "BlockQuote",
	"body":       "Body",
	"br":         "Break",
	"button":     "Button",
	"canvas":     "Canvas",
	"caption":    "Caption",
	"cite":       "Citation",
	"code":       "Code",
	"col":        "Column",
	"colgroup":   "ColumnGroup",
	"data":       "Data",
	"datalist":   "DataList",
	"dd":         "Description",
	"del":        "DeletedText",
	"details":    "Details",
	"dfn":        "Definition",
	"dialog":     "Dialog",
	"div":        "Div",
	"dl":         "DescriptionList",
	"dt":         "DefinitionTerm",
	"em":         "Emphasis",
	"embed":      "Embed",
	"fieldset":   "FieldSet",
	"figcaption": "FigureCaption",
	"figure":     "Figure",
	"footer":     "Footer",
	"form":       "Form",
	"h1":         "Heading1",
	"h2":         "Heading2",
	"h3":         "Heading3",
	"h4":         "Heading4",
	"h5":         "Heading5",
	"h6":         "Heading6",
	"header":     "Header",
	"hgroup":     "HeadingsGroup",
	"hr":         "HorizontalRule",
	"i":          "Italic",
	"iframe":     "InlineFrame",
	"img":        "Image",
	"input":      "Input",
	"ins":        "InsertedText",
	"kbd":        "KeyboardInput",
	"label":      "Label",
	"legend":     "Legend",
	"li":         "ListItem",
	"link":       "Link",
	"main":       "Main",
	"map":        "Map",
	"mark":       "Mark",
	"meta":       "Meta",
	"meter":      "Meter",
	"nav":        "Navigation",
	"noscript":   "NoScript",
	"object":     "Object",
	"ol":         "OrderedList",
	"optgroup":   "OptionsGroup",
	"option":     "Option",
	"output":     "Output",
	"p":          "Paragraph",
	"param":      "Parameter",
	"picture":    "Picture",
	"pre":        "Preformatted",
	"progress":   "Progress",
	"q":          "Quote",
	"rp":         "RubyParenthesis",
	"rt":         "RubyText",
	"rtc":        "RubyTextContainer",
	"ruby":       "Ruby",
	"s":          "Strikethrough",
	"samp":       "Sample",
	"script":     "Script",
	"section":    "Section",
	"select":     "Select",
	"slot":       "Slot",
	"small":      "Small",
	"source":     "Source",
	"span":       "Span",
	"strong":     "Strong",
	"style":      "Style",
	"sub":        "Subscript",
	"summary":    "Summary",
	"sup":        "Superscript",
	"table":      "Table",
	"tbody":      "TableBody",
	"td":         "TableData",
	"template":   "Template",
	"textarea":   "TextArea",
	"tfoot":      "TableFoot",
	"th":         "TableHeader",
	"thead":      "TableHead",
	"time":       "Time",
	"title":      "Title",
	"tr":         "TableRow",
	"track":      "Track",
	"u":          "Underline",
	"ul":         "UnorderedList",
	"var":        "Variable",
	"video":      "Video",
	"wbr":        "WordBreakOpportunity",
}

var typeProps = map[string]string{
	"button":         "TypeButton",
	"checkbox":       "TypeCheckbox",
	"color":          "TypeColor",
	"date":           "TypeDate",
	"datetime":       "TypeDatetime",
	"datetime-local": "TypeDatetimeLocal",
	"email":          "TypeEmail",
	"file":           "TypeFile",
	"hidden":         "TypeHidden",
	"image":          "TypeImage",
	"month":          "TypeMonth",
	"number":         "TypeNumber",
	"password":       "TypePassword",
	"radio":          "TypeRadio",
	"range":          "TypeRange",
	"min":            "TypeMin",
	"max":            "TypeMax",
	"value":          "TypeValue",
	"step":           "TypeStep",
	"reset":          "TypeReset",
	"search":         "TypeSearch",
	"submit":         "TypeSubmit",
	"tel":            "TypeTel",
	"text":           "TypeText",
	"time":           "TypeTime",
	"url":            "TypeUrl",
	"week":           "TypeWeek",
}

var boolProps = map[string]string{
	"autofocus": "Autofocus",
	"checked":   "Checked",
}

var stringProps = map[string]string{
	"for":         "For",
	"href":        "Href",
	"id":          "ID",
	"placeholder": "Placeholder",
	"src":         "Src",
	"value":       "Value",
}
