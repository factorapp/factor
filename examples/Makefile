.PHONY: wasm
wasm: 
	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o ./app/example.wasm ./client

.PHONY: server-app
server-app:
	go build -o server-app ./server

.PHONY: run
run: wasm server-app
	./server-app
