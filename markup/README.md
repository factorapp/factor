# markup
[![Build Status](https://travis-ci.org/murlokswarm/markup.svg?branch=master)](https://travis-ci.org/murlokswarm/markup)
[![Go Report Card](https://goreportcard.com/badge/github.com/murlokswarm/markup)](https://goreportcard.com/report/github.com/murlokswarm/markup)
[![Coverage Status](https://coveralls.io/repos/github/murlokswarm/markup/badge.svg?branch=master)](https://coveralls.io/github/murlokswarm/markup?branch=master)
[![GoDoc](https://godoc.org/github.com/murlokswarm/markup?status.svg)](https://godoc.org/github.com/murlokswarm/markup)

Package markup implements a markup language to build user interfaces.

Markups are based on HTML. They must be declared in the Render method when
implementing the Componer interface.
A markup must follow these rules:
- Regular HTML elements must be in lowercase.
- Root element of a component must be a standard HTML tag.
- Component element must have its first letter capitalized.
- Component element attribute must have its first letter capitalized.
- Each element must have a closing tag (as in XHTML).
- HTML event handlers should start with '_'.
- Template must follow the rules of https://golang.org/pkg/text/template.

## Examples
Hello component:
```go 
type Hello struct {
	Name string
}

func (c *Hello) OnInputChange(v string) string {
	c.Name = v
	app.Render(c)
}

func (c *Hello) Render() string {
	return `
 <p>
  	Hello,
 	<input onchange="OnInputChange" />
 	<World Name="{{.Name}}" />
 </p>
 	`
}
```

World component:
```go 
type World struct {
	Name string
}

func (c *World) Render() string {
	return `
 <span>
 	{{if len .Name}}
    	{{.Name}}
  	{{else}}
      	World
  	{{end}}
 </span>
 	`
}

```