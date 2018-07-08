package codegen

// NavComponentGo returns the Go code for the Nav component
//
// TODO: make this flexible to generate arbitrary components
// see https://github.com/factorapp/factor/issues/24
const NavComponentGo = `package components

import (
	"github.com/gowasm/vecty"
)

type Nav struct {
	vecty.Core
	MyProp      string ` + "`" + "vecty:" + `"Prop"
	CurrentPath string
}
`

// NavComponentHTML returns the HTML code for the Nav component
const NavComponentHTML = `<nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
<a class="navbar-brand" href="#">Navbar</a>
<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
  <span class="navbar-toggler-icon"></span>
</button>

<div class="collapse navbar-collapse" id="navbarsExampleDefault">
  <ul class="navbar-nav mr-auto">
	<li class="nav-item active">
	  <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
	</li>
	<li class="nav-item">
	  <a class="nav-link" href="#">Link</a>
	</li>
	<li class="nav-item">
	  <a class="nav-link disabled" href="#">Disabled</a>
	</li>
	<li class="nav-item dropdown">
	  <a class="nav-link dropdown-toggle" href="https://example.com" id="dropdown01" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Dropdown</a>
	  <div class="dropdown-menu" aria-labelledby="dropdown01">
		<a class="dropdown-item" href="#">Action</a>
		<a class="dropdown-item" href="#">Another action</a>
		<a class="dropdown-item" href="#">Something else here</a>
	  </div>
	</li>
  </ul>
  <form class="form-inline my-2 my-lg-0">
	<input class="form-control mr-sm-2" type="text" placeholder="Search" aria-label="Search"/>
	<button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
  </form>
</div>
</nav>`
