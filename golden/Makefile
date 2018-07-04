wasm: 
	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm ./client 
	mv example.wasm ./app/

run:  wasm
	cd server && go build && cd .. && ./server/server