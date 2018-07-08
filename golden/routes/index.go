package routes

import (
	"fmt"

	"github.com/gowasm/vecty"
)

type Index struct {
	vecty.Core
}

func (i *Index) OnClick(e *vecty.Event) {
	fmt.Println("Someone clicked on me", e.Target)
}
