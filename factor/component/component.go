package component

import "io"

// A Component represents a web component in an HTML file
type Component struct {
	Name     string
	Template string
	Style    string
}

// FromFile parses a component file like Nav.html into
// a Component structure
func FromFile(r io.Reader) (*Component, error) {

	return &Component{}, nil
}
