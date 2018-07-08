Humble/detect
=============

[![Version](https://img.shields.io/badge/version-0.1.2-5272B4.svg)](https://github.com/go-humble/detect/releases)
[![GoDoc](https://godoc.org/github.com/go-humble/detect?status.svg)](https://godoc.org/github.com/go-humble/detect)

Detect is a tiny go package for detecting whether code is running on the server
or browser. It is intended to be used with hybrid go applications and
[gopherjs](https://github.com/gopherjs/gopherjs). Detect works great as a
stand-alone package or in combination with other
[Humble](https://github.com/go-humble) packages.

Detect is written in pure go. It feels like go, follows go idioms when possible, and
compiles with the go tools. Detect runs on the server and can also be compiled
to javascript and run in the browser.


Browser Support
---------------

Detect works with IE9+ (with a
[polyfill for typed arrays](https://github.com/inexorabletash/polyfill/blob/master/typedarray.js))
and all other modern browsers. Detect compiles to javascript via [gopherjs](https://github.com/gopherjs/gopherjs)
and this is a gopherjs limitation.


Installation
------------

Install detect like you would any other go package:

```bash
go get github.com/go-humble/detect
```

You will also need to install gopherjs if you don't already have it. The latest version is
recommended. Install gopherjs with:

```
go get -u github.com/gopherjs/gopherjs
```

You can compile your application to javascript using the `gopherjs build` command. Run
`gopherjs --help` to learn more about the gopherjs command-line tool.


Example Usage
-------------

Detect is intended to be used in hybrid go applications which use gopherjs to
compile to javascript. When you are sharing code between the server and browser,
it is often useful to be able to quickly detect which platform the code is
currently running on. For example, if you are sharing models between the browser
and server, you probably want the server-side code to communicate with the
database, whereas the browser-side code should not.

```go
// The Todo type and its fields can be shared between the server and browser.
type Todo struct {
   Title       string
   IsCompleted bool
}

// The Toggle method can be called server-side or browser-side because it does
// not touch the database.
func (t *Todo) Toggle() {
   t.IsCompleted = !t.IsCompleted
}

// The Save method however, should have different behavior depending on whether
// we are running on the server or browser.
func (t *Todo) Save() {
   switch {
   case detect.IsBrowser():
      // In this case, we might want to save to local storage first, and then
      // try to sync with the server using WebSockets or http.
   case detect.IsServer():
      // In this case, we want to save directly to the database.
   }
}
```

How it Works
------------

To see if code is running as javascript, detect checks whether `js.Global`
is nil. `js.Global` is a special gopherjs property that is only non-nil when
code is compiled to javascript.

To see if code is running in the browser, detect checks for the existence of a
global javascript "document" property. If it is defined, detect assumes the
code is running in the browser.

We encourage you to check out the
[source code](https://github.com/go-humble/detect/blob/master/detect.go) if
you're curious about how detect works under the hood. It's really tiny
(only a single file) and pretty straightforward.


Contributing
------------

See [CONTRIBUTING.md](https://github.com/go-humble/detect/blob/master/CONTRIBUTING.md)


License
-------

Detect is licensed under the MIT License. See the [LICENSE](https://github.com/go-humble/detect/blob/master/LICENSE)
file for more information.
