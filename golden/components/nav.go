package components

import (
	"github.com/bketelsen/factor/markup"
)

type Nav struct {
	MyProp      string
	CurrentPath string
}

func (n *Nav) OnAfterPrint(e *markup.Event) {

}
