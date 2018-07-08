package codegen

// Makefile returns the Makefile code for the new app
//
// TODO: this won't be necessary after https://github.com/factorapp/factor/issues/11
const Makefile = `wasm: 
	GOARCH=wasm GOOS=js go build -o example.wasm ./client 
	mv example.wasm ./app/

run:  wasm
	cd server && go build && cd .. && ./server/server`
