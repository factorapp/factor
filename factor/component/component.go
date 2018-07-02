package component

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

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
		parsed:   true,
	}
	return c, err
}

func (c *Component) Transform() error {
	if !c.parsed {
		return ErrComponentNotParsed
	}
	tpl := template.Must(template.New("component").Parse(comptpl))
	err := tpl.Execute(os.Stdout, c)
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
