// Copyright 2015 Alex Browne.  All rights reserved.
// Use of this source code is governed by the MIT
// license, which can be found in the LICENSE file.

// Package detect is a tiny package for detecting whether code is running on
// the server or browser. It is intended to be used with hybrid go applications
// and gopherjs.
//
// Version 0.1.2
package detect

import (
	"github.com/gopherjs/gopherjs/js"
)

// IsBrowser returns true iff the code that is currently running is compiled
// javascript code running in a browser. It works by checking if the global
// "document" property exists.
func IsBrowser() bool {
	return IsJavascript() && js.Global.Get("document") != js.Undefined
}

// IsServer returns true iff the code is running on a server. It returns true
// if the code is pure go, or if it has been compiled to javascript and is
// being run with node. It will only return false if the global "document"
// property exists.
func IsServer() bool {
	if IsJavascript() {
		return !IsBrowser()
	}
	return true
}

// IsJavascript return true iff the code that is currently running is
// transpiled javascript code. It works by checking if the js.Global property
// is non-nil.
func IsJavascript() bool {
	return js.Global != nil
}

// IsGo returns true iff the code that is currently running is pure go code.
// That is, it has not been compiled to javascript. It works by checking if the
// js.Global property is nil.
func IsGo() bool {
	return !IsJavascript()
}
