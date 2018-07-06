// build +js,wasm
package main

import (
	"fmt"
	"log"

	"github.com/bketelsen/factor/golden/components"
	_ "github.com/bketelsen/factor/golden/routes"
	"github.com/bketelsen/factor/markup"
)

func main() {
	c := &components.App{}
	node, err := markup.MountBody(c)
	fmt.Println(node)
	if err != nil {
		log.Fatal(err)
	}
}