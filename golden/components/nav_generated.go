// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package components

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

func (p *Nav) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Navigation(
			vecty.Markup(
				vecty.Class("navbar", "navbar-expand-md", "navbar-dark", "bg-dark", "fixed-top"),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("navbar-brand"),
					prop.Href("#"),
				),
				vecty.Text("Navbar"),
			),
			elem.Button(
				vecty.Markup(
					vecty.Class("navbar-toggler"),
					prop.Type(prop.TypeButton),
					vecty.Data("toggle", "collapse"),
					vecty.Data("target", "#navbarsExampleDefault"),
					vecty.Attribute("aria-controls", "navbarsExampleDefault"),
					vecty.Attribute("aria-expanded", "false"),
					vecty.Attribute("aria-label", "Toggle navigation"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("navbar-toggler-icon"),
					),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("collapse", "navbar-collapse"),
					prop.ID("navbarsExampleDefault"),
				),
				elem.UnorderedList(
					vecty.Markup(
						vecty.Class("navbar-nav", "mr-auto"),
					),
					elem.ListItem(
						vecty.Markup(
							vecty.Class("nav-item", "active"),
						),
						elem.Anchor(
							vecty.Markup(
								vecty.Class("nav-link"),
								prop.Href("#"),
							),
							vecty.Text("Home"),
							elem.Span(
								vecty.Markup(
									vecty.Class("sr-only"),
								),
								vecty.Text("(current)"),
							),
						),
					),
					elem.ListItem(
						vecty.Markup(
							vecty.Class("nav-item"),
						),
						elem.Anchor(
							vecty.Markup(
								vecty.Class("nav-link"),
								prop.Href("#"),
							),
							vecty.Text("Link"),
						),
					),
					elem.ListItem(
						vecty.Markup(
							vecty.Class("nav-item"),
						),
						elem.Anchor(
							vecty.Markup(
								vecty.Class("nav-link", "disabled"),
								prop.Href("#"),
							),
							vecty.Text("Disabled"),
						),
					),
					elem.ListItem(
						vecty.Markup(
							vecty.Class("nav-item", "dropdown"),
						),
						elem.Anchor(
							vecty.Markup(
								vecty.Class("nav-link", "dropdown-toggle"),
								prop.Href("https://example.com"),
								prop.ID("dropdown01"),
								vecty.Data("toggle", "dropdown"),
								vecty.Attribute("aria-haspopup", "true"),
								vecty.Attribute("aria-expanded", "false"),
							),
							vecty.Text("Dropdown"),
						),
						elem.Div(
							vecty.Markup(
								vecty.Class("dropdown-menu"),
								vecty.Attribute("aria-labelledby", "dropdown01"),
							),
							elem.Anchor(
								vecty.Markup(
									vecty.Class("dropdown-item"),
									prop.Href("#"),
								),
								vecty.Text("Action"),
							),
							elem.Anchor(
								vecty.Markup(
									vecty.Class("dropdown-item"),
									prop.Href("#"),
								),
								vecty.Text("Another action"),
							),
							elem.Anchor(
								vecty.Markup(
									vecty.Class("dropdown-item"),
									prop.Href("#"),
								),
								vecty.Text("Something else here"),
							),
						),
					),
				),
				elem.Form(
					vecty.Markup(
						vecty.Class("form-inline", "my-2", "my-lg-0"),
					),
					elem.Input(
						vecty.Markup(
							vecty.Class("form-control", "mr-sm-2"),
							prop.Type(prop.TypeText),
							prop.Placeholder("Search"),
							vecty.Attribute("aria-label", "Search"),
						),
					),
					elem.Button(
						vecty.Markup(
							vecty.Class("btn", "btn-outline-success", "my-2", "my-sm-0"),
							prop.Type(prop.TypeSubmit),
						),
						vecty.Text("Search"),
					),
				),
			),
		),
	)
}
