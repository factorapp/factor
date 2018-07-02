package component

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

// A Component represents a web component in an HTML file
type Component struct {
	Name     string
	Template string
	Style    string
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
	}
	return c, err
}

func removeStyleTags(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "<style>", "", -1)
	s = strings.Replace(s, "</style>", "", -1)
	return s
}
