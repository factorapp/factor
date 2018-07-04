// build +js,wasm
package main

import (
	"log"

	"github.com/bketelsen/factor/golden/components"
	"github.com/bketelsen/factor/markup"
)

func main() {

	c := &components.App{}
	_, err := markup.MountBody(c)
	if err != nil {
		log.Fatal(err)
	}
}
