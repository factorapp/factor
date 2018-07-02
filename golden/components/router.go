package components

// import "github.com/bketelsen/factor/markup"

// type routeMaker = func() (markup.Componer, error)

// func newRouteMaker(tag string) routeMaker {
// 	return func() (markup.Componer, error) {
// 		return markup.New(tag)
// 	}
// }

// var routes = map[string]routeMaker{
// 	"/about", newRouteMaker("routes.About"),
// 	"/blog/{slug}": newRouteMaker("routes.Blog"),
// }

// type Router struct {
// 	routes map[string]routeMaker
// }

// func NewRouter() Router {
// 	return Router{routes: routes}
// }

// func (r *Router) Route(path string) markup.Componer {
// 	fn, ok := r.routes[path]
// 	if !ok {
// 		return ForOhForComp{}
// 	}
// 	comp, err := fn()
// 	if err != nil {
// 		return ForOhForComp{}
// 	}
// 	return comp
// }

// func (r *Router) Calculate() error {

// 	// walk the routes directory

// 	// add a routeComponent entry for each html/go combo

// 	// something with regex

// 	// Profit
// }
