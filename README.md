# Factor

Factor is a tool for building SPA web applications in Go using Web Assembly (WASM).

Factor is heavily inspired by [Svelte and Sapper](https://sapper.svelte.technology/guide#getting-started), but doesn't attempt to be a straight port, but rather take those ideas and present them in a way natural for a Go developer.

## Application Structure

```
├ app
│ ├ App.html # Main/initial view
│ └ template.html # application template
├ assets
│ ├ # your files here
├ components 
│ ├ # your files here
├ routes
│ ├ # your routes here - these are pages.
│ ├ _error.html
│ └ index.html
```

## Requirements

WASM support is only available in Go 1.11, which is not yet officially released, so you'll need to build it from source. Full instructions on doing so are [here](https://golang.org/doc/install/source), but we're making it easy on you here:

```console
$ git clone https://go.googlesource.com/go
$ cd go
$ git checkout v1.11beta1
$ cd src
$ ./all.bash
```

Once you see `ALL TESTS PASSED`, your Go 1.11 beta 1 toolchain is done! Now, move
the entire `go` directory to `~/gowasm`:

```console
$ cd ..
$ mv go ~/gowasm
```

And you're good to go :)

## Development

First, make sure you have built your Go 1.11 beta 1 toolchain (see "Requirements" above). Then, do this:

```console
$ cd factor
$ go install
```

Make sure you have `$GOPATH/bin` in your executable path (i.e. `$PATH`) and your `factor` CLI is ready to go, hot off the presses.

## Notes

* Factor enforces very strict HTML rules.  `<br>` will break.  Close all self-closing tags properly: `<br />`.  Failing to do this will result in an error like `element input closed by form` in your generated Go files.
