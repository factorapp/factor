package main

import (
	"log"
	"net/http"

	"github.com/bketelsen/factor/golden/components"
	"github.com/bketelsen/factor/markup"
)

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "example.wasm")
}

func main() {
	c := &components.App{}
	_, err := markup.MountBody(c)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/example.wasm", wasmHandler)
	log.Fatal(http.ListenAndServe(":3000", mux))

}
