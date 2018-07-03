package main

import (
	"log"
	"net/http"
)

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "example.wasm")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/example.wasm", wasmHandler)
	log.Printf("serving on :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
