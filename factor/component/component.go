package component

import (
	"bytes"
	"io"
	"io/ioutil"
	"math/rand"
	"strings"
	"text/template"
	"time"

	"github.com/tdewolff/parse"
	"github.com/tdewolff/parse/css"

	"github.com/pkg/errors"
)

// ErrComponentNotParsed is returned when an attempt is made
// to Transform() a Component before calling Parse()
var ErrComponentNotParsed = errors.New("transform must be called after parse")

// A Component represents a web component in an HTML file
type Component struct {
	Name     string
	Template string
	Style    string
	Package  string
	Imports  []string
	parsed   bool
	Struct   bool
	UniqueID string
}

func (c *Component) WriteImports() string {
	return strings.Join(c.Imports, "\n\t")
}
func (c *Component) QuotedStyle() string {
	return "`" + c.Style + "`"
}

func (c *Component) QuotedTemplate() string {
	return "`" + c.Template + "`"
}

func (c *Component) TransformStyle() {
	p := css.NewParser(bytes.NewBufferString(c.Style), false)

	output := ""
	for {
		grammar, _, data := p.Next()
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
				data = append(data, c.UniqueID...)
				data = append(data, "{"...)
			} else if grammar == css.AtRuleGrammar || grammar == css.DeclarationGrammar || grammar == css.CustomPropertyGrammar {
				data = append(data, ";"...)
			} else if grammar == css.QualifiedRuleGrammar {
				data = append(data, ","...)
			}
		}
		output += string(data)
	}

	c.Style = output

}

// Parse reads a component file like Nav.html into
// a Component structure
func Parse(r io.Reader, name string) (*Component, error) {
	var template, style string
	var err error

	// TODO: do this more cleanly, not the fast/hacky way

	bb, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "reading component template")
	}

	s := string(bb)
	styleStart := strings.Index(s, "<style>")
	if styleStart == -1 {
		styleStart = len(s)
	}
	template = strings.TrimSpace(s[:styleStart])
	if styleStart != len(s) {
		style = strings.TrimSpace(s[styleStart:])
		style = strings.Replace(style, "<style>", "", -1)
		style = strings.Replace(style, "</style>", "", -1)
	}
	c := &Component{
		Name:     name,
		Template: template,
		Style:    style,
		UniqueID: randSeq(10),
		parsed:   true,
	}
	return c, err
}

func (c *Component) Transform(w io.Writer) error {
	if !c.parsed {
		return ErrComponentNotParsed
	}
	tpl := template.Must(template.New("component").Parse(comptpl))
	err := tpl.Execute(w, c)
	if err != nil {
		return err
	}
	return nil
}

func removeStyleTags(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "<style>", "", -1)
	s = strings.Replace(s, "</style>", "", -1)
	return s
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
