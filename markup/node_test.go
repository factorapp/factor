package markup

import (
	"testing"

	"github.com/satori/go.uuid"
)

func TestNodeString(t *testing.T) {
	t.Log(Node{
		ID:  uuid.Must(uuid.NewV1()),
		Tag: "div",
	})
}

func TestNodeMarkup(t *testing.T) {
	n := Node{
		ID:  uuid.Must(uuid.NewV1()),
		Tag: "a",
		Attributes: AttributeMap{
			"href": "Hello",
		},
	}
	t.Log(n.Markup())
}

func TestNodeMarkupHrefError(t *testing.T) {
	n := Node{
		ID:  uuid.Must(uuid.NewV1()),
		Tag: "a",
		Attributes: AttributeMap{
			"href": "% sga -= Coucou maman",
		},
	}
	t.Log(n.Markup())
}
