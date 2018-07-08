// Copyright 2015 Alex Browne and Soroush Pour.
// Allrights reserved. Use of this source code is
// governed by the MIT license, which can be found
// in the LICENSE file.

package router

import (
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/go-humble/detect"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var (
	// browserSupportsPushState will be true if the current browser
	// supports history.pushState and the onpopstate event.
	browserSupportsPushState bool
	document                 dom.HTMLDocument
)

func init() {
	if detect.IsBrowser() {
		// We only want to initialize certain things if we are running
		// inside a browser. Otherwise, they will cause the program to
		// panic.
		var ok bool
		document, ok = dom.GetWindow().Document().(dom.HTMLDocument)
		if !ok {
			panic("Could not convert document to dom.HTMLDocument")
		}
		browserSupportsPushState = (js.Global.Get("onpopstate") != js.Undefined) &&
			(js.Global.Get("history") != js.Undefined) &&
			(js.Global.Get("history").Get("pushState") != js.Undefined)
	}
}

// Router is responsible for handling routes. If history.pushState is
// supported, it uses it to navigate from page to page and will listen
// to the "onpopstate" event. Otherwise, it sets the hash component of the
// url and listens to changes via the "onhashchange" event.
type Router struct {
	// routes is the set of routes for this router.
	routes []*route
	// ShouldInterceptLinks tells the router whether or not to intercept click events
	// on links and call the Navigate method instead of the default behavior.
	// If it is set to true, the router will automatically intercept links when
	// Start, Navigate, or Back are called, or when the onpopstate event is triggered.
	ShouldInterceptLinks bool
	// ForceHashURL tells the router to use the hash component of the url to
	// represent different routes, even if history.pushState is supported.
	ForceHashURL bool
	// Verbose determines whether or not the router will log to console.log.
	// If true, the router will log a message if, e.g., a match cannot be found for
	// a particular path.
	Verbose bool
	// listener is the js.Object representation of a listener callback.
	// It is required in order to use the RemoveEventListener method
	listener func(*js.Object)
}

// Context is used as an argument to Handlers
type Context struct {
	// Params is the parameters from the url as a map of names to values.
	Params map[string]string
	// Path is the path that triggered this particular route. If the hash
	// fallback is being used, the value of path does not include the '#'
	// symbol.
	Path string
	// InitialLoad is true iff this route was triggered during the initial
	// page load. I.e. it is true if this is the first path that the browser
	// was visiting when the javascript finished loading.
	InitialLoad bool
	// QueryParams is the query params from the URL. Because params may be
	// repeated with different values, the value part of the map is a slice
	QueryParams map[string][]string
}

// Handler is a function which is run in response to a specific
// route. A Handler takes a Context as an argument, which gives
// handler functions access to path parameters and other important
// information.
type Handler func(context *Context)

// New creates and returns a new router
func New() *Router {
	return &Router{
		routes: []*route{},
	}
}

// route is a representation of a specific route
type route struct {
	// regex is a regex pattern that matches route
	regex *regexp.Regexp
	// paramNames is an ordered list of parameters expected
	// by route handler
	paramNames []string
	// handler called when route is matched
	handler Handler
}

// HandleFunc will cause the router to call f whenever window.location.pathname
// (or window.location.hash, if history.pushState is not supported) matches path.
// path can contain any number of parameters which are denoted with curly brackets.
// So, for example, a path argument of "users/{id}" will be triggered when the user
// visits users/123 and will call the handler function with params["id"] = "123".
func (r *Router) HandleFunc(path string, handler Handler) {
	r.routes = append(r.routes, newRoute(path, handler))
}

// newRoute returns a route with the given arguments. paramNames and regex
// are calculated from the path
func newRoute(path string, handler Handler) *route {
	route := &route{
		handler: handler,
	}
	strs := strings.Split(path, "/")
	strs = removeEmptyStrings(strs)
	pattern := `^`
	for _, str := range strs {
		if str[0] == '{' && str[len(str)-1] == '}' {
			pattern += `/`
			pattern += `([^/]*)`
			route.paramNames = append(route.paramNames, str[1:(len(str)-1)])
		} else {
			pattern += `/`
			pattern += str
		}
	}
	pattern += `/?$`
	route.regex = regexp.MustCompile(pattern)
	return route
}

// Start causes the router to listen for changes to window.location and
// trigger the appropriate handler whenever there is a change.
func (r *Router) Start() {
	if browserSupportsPushState && !r.ForceHashURL {
		r.pathChanged(getPath(), true)
		r.watchHistory()
	} else {
		r.setInitialHash()
		r.watchHash()
	}
	if r.ShouldInterceptLinks {
		r.InterceptLinks()
	}
}

// Stop causes the router to stop listening for changes, and therefore
// the router will not trigger any more router.Handler functions.
func (r *Router) Stop() {
	if browserSupportsPushState && !r.ForceHashURL {
		js.Global.Set("onpopstate", nil)
	} else {
		js.Global.Set("onhashchange", nil)
	}
}

// Navigate will trigger the handler associated with the given path
// and update window.location accordingly. If the browser supports
// history.pushState, that will be used. Otherwise, Navigate will
// set the hash component of window.location to the given path.
func (r *Router) Navigate(path string) {
	if browserSupportsPushState && !r.ForceHashURL {
		pushState(path)
		r.pathChanged(path, false)
	} else {
		setHash(path)
	}
	if r.ShouldInterceptLinks {
		r.InterceptLinks()
	}
}

// CanNavigate returns true if the specified path can be navigated by the
// router, and false otherwise
func (r *Router) CanNavigate(path string) bool {
	if bestRoute, _, _ := r.findBestRoute(path); bestRoute != nil {
		return true
	}
	return false
}

// Back will cause the browser to go back to the previous page.
// It has the same effect as the user pressing the back button,
// and is just a wrapper around history.back()
func (r *Router) Back() {
	js.Global.Get("history").Call("back")
	if r.ShouldInterceptLinks {
		r.InterceptLinks()
	}
}

// InterceptLinks intercepts click events on links of the form <a href="/foo"></a>
// and calls router.Navigate("/foo") instead, which triggers the appropriate Handler
// instead of requesting a new page from the server. Since InterceptLinks works by
// setting event listeners in the DOM, you must call this function whenever the DOM
// is changed. Alternatively, you can set r.ShouldInterceptLinks to true, which will
// trigger this function whenever Start, Navigate, or Back are called, or when the
// onpopstate event is triggered. Even with r.ShouldInterceptLinks set to true, you
// may still need to call this function if you change the DOM manually without
// triggering a route.
func (r *Router) InterceptLinks() {
	for _, link := range document.Links() {
		href := link.GetAttribute("href")
		switch {
		case href == "":
			continue

		case strings.HasPrefix(href, "http://"), strings.HasPrefix(href, "https://"), strings.HasPrefix(href, "//"):
			// These are external links and should behave normally.
			continue

		case strings.HasPrefix(href, "#"):
			// These are anchor links and should behave normally.
			// Recall that even when we are using the hash trick, href
			// attributes should be relative paths without the "#" and
			// router will handle them appropriately.
			continue

		case strings.HasPrefix(href, "/"):
			// These are relative links. The kind that we want to intercept.
			if r.listener != nil {
				// Remove the old listener (if any)
				link.RemoveEventListener("click", true, r.listener)
			}
			r.listener = link.AddEventListener("click", true, r.interceptLink)
		}
	}
}

// interceptLink is intended to be used as a callback function. It stops
// the default behavior of event and instead calls r.Navigate, passing through
// the link's href property.
func (r *Router) interceptLink(event dom.Event) {
	path := event.CurrentTarget().GetAttribute("href")
	// Only intercept the click event if we have a route which matches
	// Otherwise, just do the default.
	if bestRoute, _, _ := r.findBestRoute(path); bestRoute != nil {
		event.PreventDefault()
		go r.Navigate(path)
	}
}

// setInitialHash will set hash to / if there is currently no hash.
func (r *Router) setInitialHash() {
	if getHash() == "" {
		setHash("/")
	} else {
		r.pathChanged(getPathFromHash(getHash()), true)
	}
}

// pathChanged should be called whenever the path changes and will trigger
// the appropriate handler. initial should be true iff this is the first
// time the javascript is loaded on the page.
func (r *Router) pathChanged(path string, initial bool) {
	bestRoute, tokens, params := r.findBestRoute(path)
	// If no routes match, we throw console error and no handlers are called
	if bestRoute == nil {
		if r.Verbose {
			log.Println("Could not find route to match: " + path)
		}
		return
	}
	// Create the context and pass it through to the handler
	c := &Context{
		Path:        path,
		InitialLoad: initial,
		Params:      map[string]string{},
		QueryParams: params,
	}
	for i, token := range tokens {
		c.Params[bestRoute.paramNames[i]] = token
	}
	bestRoute.handler(c)
}

// findBestRoute compares the given path against regex patterns of routes.
// Preference given to routes with most literal (non-parameter) matches. For
// example if we have the following:
//   Route 1: /todos/work
//   Route 2: /todos/{category}
// And the path argument is "/todos/work", the bestRoute would be todos/work
// because the string "work" matches the literal in Route 1.
func (r Router) findBestRoute(path string) (bestRoute *route, tokens []string, params map[string][]string) {
	parts := strings.SplitN(path, "?", 2)
	leastParams := -1
	for _, route := range r.routes {
		matches := route.regex.FindStringSubmatch(parts[0])
		if matches != nil {
			if (leastParams == -1) || (len(matches) < leastParams) {
				leastParams = len(matches)
				bestRoute = route
				tokens = matches[1:]
			}
		}
	}
	if len(parts) > 1 {
		params = r.parseQueryPart(parts[1])
	}
	return bestRoute, tokens, params
}

// parseQueryPart extracts query params from the query part of the URL
func (r Router) parseQueryPart(queryPart string) (params map[string][]string) {
	var err error
	params, err = url.ParseQuery(queryPart)
	if err != nil && r.Verbose {
		// the URL spec allows things other than name/value pairs in the query
		// part of the URL, so we optionally log a message
		log.Printf("Error parsing query %v: %v", queryPart, err)
	}
	return
}

// removeEmptyStrings removes any empty strings from strings
func removeEmptyStrings(strings []string) []string {
	result := []string{}
	for _, s := range strings {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

// watchHash listens to the onhashchange event and calls r.pathChanged when
// it changes
func (r *Router) watchHash() {
	js.Global.Set("onhashchange", func() {
		go func() {
			path := getPathFromHash(getHash())
			r.pathChanged(path, false)
		}()
	})
}

// watchHistory listens to the onpopstate event and calls r.pathChanged when
// it changes
func (r *Router) watchHistory() {
	js.Global.Set("onpopstate", func() {
		go func() {
			r.pathChanged(getPath(), false)
			if r.ShouldInterceptLinks {
				r.InterceptLinks()
			}
		}()
	})
}

// getPathFromHash returns everything after the "#" character in hash.
func getPathFromHash(hash string) string {
	return strings.SplitN(hash, "#", 2)[1]
}

// getHash is an alias for js.Global.Get("location").Get("hash").String()
func getHash() string {
	return js.Global.Get("location").Get("hash").String()
}

// setHash is an alias for js.Global.Get("location").Set("hash", hash)
func setHash(hash string) {
	js.Global.Get("location").Set("hash", hash)
}

// getPath is an alias for js.Global.Get("location").Get("pathname").String()
func getPath() string {
	return js.Global.Get("location").Get("pathname").String()
}

// pushState is an alias for js.Global.Get("history").Call("pushState", nil, "", path)
func pushState(path string) {
	js.Global.Get("history").Call("pushState", nil, "", path)
}
