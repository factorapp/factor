package markup

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"strings"

	"github.com/satori/go.uuid"
)

type decoder struct {
	xmlDecoder *xml.Decoder
	root       *Node
	current    *Node
}

func newDecoder(r io.Reader) *decoder {
	return &decoder{
		xmlDecoder: xml.NewDecoder(r),
	}
}

func (d *decoder) Decode() (root *Node, err error) {
	if err = d.next(); err != nil {
		return
	}

	if d.root == nil {
		err = errors.New("empty markup")
		return
	}

	root = d.root
	return
}

func (d *decoder) next() error {
	token, err := d.xmlDecoder.Token()
	if err != nil {
		if err == io.EOF {
			return nil
		}

		return err
	}

	switch t := token.(type) {
	case xml.StartElement:
		n := elementToNode(t)

		if d.root == nil {
			d.root = n
		}

		if d.current != nil {
			d.current.Children = append(d.current.Children, n)
		}

		n.Parent = d.current
		d.current = n

	case xml.EndElement:
		if d.current.Parent != nil {
			d.current = d.current.Parent
		}

	case xml.CharData:
		n := charDataToNode(t)

		if len(n.Text) == 0 {
			break
		}

		if d.root == nil {
			return errors.New("text nodes cannot be root")
		}

		d.current.Children = append(d.current.Children, n)
	}
	return d.next()
}

func elementToNode(e xml.StartElement) *Node {
	tag := e.Name.Local
	nodeType := HTMLNode

	if isComponentTag(tag) {
		nodeType = ComponentNode
	}

	attributes := AttributeMap{}

	for _, attr := range e.Attr {
		attributes[attr.Name.Local] = attr.Value
	}
	id, _ := uuid.NewV4()
	return &Node{
		ID:         id,
		Type:       nodeType,
		Tag:        tag,
		Attributes: attributes,
	}
}

func charDataToNode(d xml.CharData) *Node {

	id, _ := uuid.NewV4()
	return &Node{
		ID:   id,
		Type: TextNode,
		Text: strings.TrimSpace(string(d)),
	}
}

func stringToNode(v string) (root *Node, err error) {
	b := bytes.NewBufferString(v)
	dec := newDecoder(b)
	return dec.Decode()
}
