// Package router is a router for client-side web applications written in pure
// go which compiles to javascript via gopherjs
// (https://github.com/gopherjs/gopherjs).
//
// Version 0.5.0
//
// Router supports the following features:
//
//   - Write code in pure go. It feels like go, follows go idioms, and
//     compiles with the go tools.
//   - Each route consists of a path and a handler function which is
//     triggered when path matches.
//   - Routes can have parameters, which are passed through as an
//     argument to handler functions.
//   - Router uses history.pushState and listens to the onpopstate
//     event in browsers that support it. In older browsers it
//     automatically falls back to using a hash.
//   - Router can be configured to automatically intercept link click
//     events, triggering the appropriate route instead of requesting
//     a new page.
//
// For the full source code, a getting started guide, and more information visit
// https://github.com/go-humble/router.
package router
