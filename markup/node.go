package markup

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"strings"
	"syscall/js"

	"github.com/satori/go.uuid"
	"github.com/tdewolff/parse"
	"github.com/tdewolff/parse/css"
)

// Enumeration of the node types.
const (
	HTMLNode NodeType = iota
	ComponentNode
	TextNode
)

var (
	selfClosingTags = map[string]bool{
		"area":    true,
		"base":    true,
		"br":      true,
		"col":     true,
		"command": true,
		"embed":   true,
		"hr":      true,
		"img":     true,
		"input":   true,
		"link":    true,
		"meta":    true,
		"param":   true,
		"source":  true,
	}
)

// Node represents a markup node.
type Node struct {
	ID             uuid.UUID
	ContextID      uuid.UUID
	Type           NodeType
	Tag            string
	Text           string
	Attributes     AttributeMap
	Component      Componer
	Mount          Componer
	Parent         *Node
	Children       []*Node
	Element        js.Value
	eventListeners []*EventListener
}

// NodeType represents the type of the node.
type NodeType uint8

//  String returns a string representing the node.
func (n *Node) String() string {
	return fmt.Sprintf("[\033[36m%v\033[00m \033[33m%v\033[00m]", n.Tag, n.ID)
}

// Markup return a string which contains the markup of the node.
func (n *Node) Markup() string {
	return n.markup(0)
}
func (n *Node) transformStyle(style string) string {
	p := css.NewParser(bytes.NewBufferString(style), false)

	output := ""
	for {
		grammar, _, data := p.Next()
		fmt.Println(grammar, string(data))
		data = parse.Copy(data)
		if grammar == css.ErrorGrammar {
			if err := p.Err(); err != io.EOF {
				for _, val := range p.Values() {
					data = append(data, val.Data...)
				}
				if perr, ok := err.(*parse.Error); ok && perr.Message == "unexpected token in declaration" {
					data = append(data, ";"...)
				}
			} else {
				break
			}
		} else if grammar == css.AtRuleGrammar || grammar == css.BeginAtRuleGrammar || grammar == css.QualifiedRuleGrammar || grammar == css.BeginRulesetGrammar || grammar == css.DeclarationGrammar || grammar == css.CustomPropertyGrammar {
			if grammar == css.DeclarationGrammar || grammar == css.CustomPropertyGrammar {
				data = append(data, ":"...)
			}
			for _, val := range p.Values() {
				data = append(data, val.Data...)
			}
			if grammar == css.BeginAtRuleGrammar || grammar == css.BeginRulesetGrammar {

				data = append(data, "."...)
				data = append(data, n.ID.String()...)
				data = append(data, "{"...)

			} else if grammar == css.EndRulesetGrammar || grammar == css.EndAtRuleGrammar {
				data = append(data, "\n"...)
			} else if grammar == css.AtRuleGrammar || grammar == css.DeclarationGrammar || grammar == css.CustomPropertyGrammar {
				data = append(data, ";"...)
			} else if grammar == css.QualifiedRuleGrammar {
				data = append(data, ","...)
			}
		}
		data = append(data, "\n"...)
		output += string(data)
	}

	return output
}
func (n *Node) markup(indent int) string {
	fmt.Println("Node", n)
	b := &bytes.Buffer{}
	indt := indentation(indent)
	fmt.Println("indent", indent)
	if n.Type == TextNode {
		b.WriteString(indt)
		b.WriteString(html.EscapeString(n.Text))
		return b.String()
	}

	if n.Type == ComponentNode {
		if n.Component == nil {
			b.WriteString(indt)
			b.WriteString("<!-- ")
			b.WriteString(n.Tag)
			b.WriteString(" -->")
			return b.String()
		}

		b.WriteString(Root(n.Component).markup(indent))
		return b.String()
	}

	b.WriteString(indt)
	b.WriteRune('<')
	b.WriteString(n.Tag)
	/*	b.WriteRune(' ')
		b.WriteString(`id="`)
		b.WriteString(n.ID.String())
		b.WriteRune('"')
	*/
	b.WriteRune(' ')
	b.WriteString(`class="`)
	b.WriteString(n.ID.String())
	b.WriteRune('"')
	for name, value := range n.Attributes {
		b.WriteRune(' ')

		if isMarkupEvent(name) {
			b.WriteString(name)
			b.WriteString(`="CallEvent('`)
			b.WriteString(n.ID.String())
			b.WriteString("', '")
			b.WriteString(value)
			b.WriteString(`', this, event)"`)

			continue
		}

		/*	if name == "href" {
				URL, err := url.Parse(value)
				if err != nil {
					log.Errorf("invalid url: %s", value)
					continue
				}

				if len(URL.Scheme) == 0 {
					URL.Scheme = "component"
				}

				b.WriteString(name)
				b.WriteString(`="`)
				b.WriteString(URL.String())
				b.WriteRune('"')
				continue
			}
		*/

		b.WriteString(name)
		b.WriteString(`="`)
		b.WriteString(value)
		b.WriteRune('"')
	}

	if _, selfClosing := selfClosingTags[n.Tag]; selfClosing {
		b.WriteString("/>")
		return b.String()
	}

	if len(n.Children) == 0 {
		b.WriteString("></")
		b.WriteString(n.Tag)
		b.WriteRune('>')
		return b.String()
	}

	b.WriteString(">\n")

	for _, child := range n.Children {
		b.WriteString(child.markup(indent + 1))
		b.WriteRune('\n')
	}

	b.WriteString(indt)
	b.WriteString("</")
	b.WriteString(n.Tag)
	b.WriteRune('>')
	return b.String()
}

func indentation(n int) string {
	b := bytes.Buffer{}

	for i := 0; i < n; i++ {
		b.WriteString("  ")
	}
	return b.String()
}

func isMarkupEvent(v string) bool {
	return strings.HasPrefix(v, "on")
}

func isComponentTag(tag string) bool {
	return len(tag) > 0 && tag[0] >= 'A' && tag[0] <= 'Z'
}
