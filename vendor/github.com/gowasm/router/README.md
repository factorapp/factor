Humble/Router
=============

[![Version](https://img.shields.io/badge/version-0.5.0-5272B4.svg)](https://github.com/go-humble/router/releases)
[![GoDoc](https://godoc.org/github.com/go-humble/router?status.svg)](https://godoc.org/github.com/go-humble/router)

A router for client-side web applications written in pure go which compiles to
javascript via [gopherjs](https://github.com/gopherjs/gopherjs). Router works
great as a stand-alone package or in combination with other
[Humble](https://github.com/go-humble/humble) packages.

Router supports the following features:

- Write code in pure go. It feels like go, follows go idioms, and compiles with the go tools.
- Each route consists of a path and a handler function which is triggered when path matches.
- Routes can have parameters, which are passed through as an argument to handler functions.
- Router uses history.pushState and listens to the onpopstate event in browsers that support
  it. In older browsers it automatically falls back to using a hash.
- Router can be configured to automatically intercept link click events, triggering the appropriate
  route instead of requesting a new page.


Browser Support
---------------

Router works with IE9+ (with a
[polyfill for typed arrays](https://github.com/inexorabletash/polyfill/blob/master/typedarray.js))
and all other modern browsers. Router compiles to javascript via [gopherjs](https://github.com/gopherjs/gopherjs)
and this is a gopherjs limitation.

Router is regularly tested with the latest versions of Firefox, Chrome, and Safari on Mac OS.
Each major or minor release is tested with IE9+ and the latest versions of Firefox and Chrome
on Windows.


Installation
------------

Install router like you would any other go package:

```bash
go get github.com/go-humble/router
```

You will also need to install gopherjs if you don't already have it. The latest
version is recommended. Install gopherjs with:

```
go get -u github.com/gopherjs/gopherjs
```

You can compile your application to javascript using the `gopherjs build`
command. Run `gopherjs --help` to learn more about the gopherjs command-line
tool.


Quickstart Guide
----------------

### Declaring Routes

Declaring routes works similarly to other routing packages in go. Here's an example:

```go
// Create a new Router object
r := router.New()
// Use HandleFunc to add routes.
r.HandleFunc("/people", indexPeople)
r.HandleFunc("/people/{id}", showPerson)
// You must call Start in order to start listening for changes
// in the url and trigger the appropriate handler function.
r.Start()
```

The second argument to HandleFunc should be a router.Handler, which has the following
definition:

```go
type Handler func(context *Context)
```

And the Context type is defined as follows:

```go
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
}
```

### Accessing Parameters

Parameters are accessed via `context.Params`. If we have a route defined
like this:

```go
r.HandleFunc("/people/{id}", showPerson)
```

And the path `/people/123` is triggered, then the parameters passed into the showPerson
function would look like the following:

```go
context.Params = map[string]string{
	"id": "123",
}
```

In our showPerson function, we could access the id parameter like so:

```go
func showPerson(context *router.Context) {
	id := context.Params["id"]
	// ...
}
```

A path can have any number of parameters, but each parameter must be defined inside slashes.
So a path like `people/{id}/{action}` is supported, but `people/{id}?action={action}` is not.
This may change in the future.

### Navigating Manually

You can trigger a route manually with the `Navigate` method:

```go
// Triger the route corresponding to "people/{id}"
r.Navigate("people/123")
```

Arguments to `Navigate` should never contain the hash symbol. The router package will detect
support for `history.pushState` and automatically fallback to using hashes if it is not supported.

You can also call the `Back` function to navigate back to the previous page.

### Intercepting Link Clicks

A `router.Router` can be configured to intercept link click events and trigger the appropriate route
with the `InterceptLinks` method. When called, the router finds links of the form `<a href="/foo"></a>`
and calls `router.Navigate("/foo")` instead, which triggers the appropriate Handler instead of requesting
a new page from the server. Since `InterceptLinks` works by setting event listeners in the DOM, you must
call this function whenever the DOM is changed.

Alternatively, you can set `router.ShouldInterceptLinks` to true, which will trigger the `InterceptLinks`
method whenever `Start`, `Navigate`, or `Back` are called, or when the `onpopstate` event is triggered.
Even with `ShouldInterceptLinks` set to true, you may still need to call `InterceptLinks` if you change
the DOM manually.


Testing
-------

Router uses three different types of tests.

### Regular Go Tests

Router can be tested like any other go package by running `go test .` in the root directory.
These tests make sure that path matching functions work correctly and that parameters are passed
through to handler functions, so they don't require access to a browser.

### Gopherjs Tests

If you have [node.js](https://nodejs.org/) installed, router can also be tested with gopherjs by
running `gopherjs test github.com/go-humble/router`. The `gopherjs test` command compiles the same
tests as above into javascript and runs them with node.js. These tests make sure that we haven't
broken compatibility with gopherjs and that the code still runs properly when it is compiled to
javascript.

### Browser Tests

The third type of tests use the [karma test runner](http://karma-runner.github.io/0.12/index.html)
to test the code running in actual browsers. It makes sure that router is able to respond to events
and work correctly with the browser history.

The browser tests require additional dependencies:

- [node.js](http://nodejs.org/) (If you didn't already install it above)
- [karma](http://karma-runner.github.io/0.12/index.html)
- [karma-qunit](https://github.com/karma-runner/karma-qunit)

Don't forget to also install the karma command line tools with `npm install -g karma-cli`.

You will also need to install a launcher for each browser you want to test with, as well as the
browsers themselves. Typically you install a karma launcher with `npm install -g karma-chrome-launcher`.
You can edit the config file at `karma/test-mac.conf.js` or create a new one (e.g. `karma/test-windows.conf.js`)
if you want to change the browsers that are tested on.

Once you have installed all the dependencies, start karma with `karma start karma/test-mac.conf.js` (or
your customized config file, if applicable). Once karma is running, you can keep it running in between tests.

Next you need to compile the test.go file to javascript so it can run in the browsers:

```
gopherjs build karma/go/router_test.go -o karma/js/router_test.js
```

Finally run the tests with `karma run karma/test-mac.conf.js` (changing the name of the config file if needed).

If you are on a unix-like operating system, you can recompile and run the tests in one go by running
the provided bash script: `./karma/test.sh`.


Contributing
------------

See [CONTRIBUTING.md](https://github.com/go-humble/router/blob/master/CONTRIBUTING.md)


License
-------

Router is licensed under the MIT License. See the [LICENSE](https://github.com/go-humble/router/blob/master/LICENSE)
file for more information.
