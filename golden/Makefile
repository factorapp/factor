.PHONY: wasm
wasm: 
	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm ./wasm

.PHONY: server-app
server-app:
	go build -o server-app .

.PHONY: run
run: wasm server-app
	./server-app
