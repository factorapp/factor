package codegen

// RoutesHTML returns the HTML code for the index route
const RoutesHTML = `<body>
<components:Nav />
<main role="main" class="container">

	<div class="starter-template">
	  <h1>Bootstrap starter template</h1>
	  <p class="lead">Use this document as a way to quickly start any new project.<br/> All you get is this text and a mostly barebones HTML document.</p>
	</div>

  </main><!-- /.container -->
</body>`

// RoutesGo returns the Go code for the index route
const RoutesGo = `package routes

import (
	"fmt"
	"strconv"

	"github.com/gowasm/vecty"
)

type Index struct {
	vecty.Core
	CountText string
	count     int
}

func (i *Index) OnClick(e *vecty.Event) {
	fmt.Println("Someone clicked on me", e.Target)
	i.count++
	i.CountText = "Click Count: " + strconv.Itoa(i.count)

	vecty.Rerender(i)
}
`
