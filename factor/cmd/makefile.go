package cmd

var makefileTemplate = `wasm: 
	GOARCH=wasm GOOS=js go build -o example.wasm ./client 
	mv example.wasm ./app/

run:  wasm
	cd server && go build && cd .. && ./server/server`
