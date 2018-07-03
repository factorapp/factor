package main

import (
	"log"
	"time"

	"github.com/bketelsen/factor/golden/components"
	"github.com/bketelsen/factor/markup"
)

func main() {
	c := components.NewApp("hello")
	c.Greeting = "hello!"
	if _, err := markup.MountBody(c); err != nil {
		log.Fatalf("1. couldn't mount body (%s)", err)
	}

	time.Sleep(1 * time.Second)
	log.Printf("it's been a second")
	c = &components.App{}
	c.Greeting = "it's been second"
	if _, err := markup.MountBody(c); err != nil {
		log.Fatalf("2. couldn't mount body (%s)", err)
	}
}
