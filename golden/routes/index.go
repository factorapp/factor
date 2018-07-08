package routes

import (
	"fmt"
	"strconv"

	"github.com/gowasm/vecty"
)

type Index struct {
	vecty.Core
	Notice string
	count  int
}

func (i *Index) OnClick(e *vecty.Event) {
	fmt.Println("Someone clicked on me", e.Target)
	i.count++
	i.Notice = "Click Count: " + strconv.Itoa(i.count)

	vecty.Rerender(i)
}
