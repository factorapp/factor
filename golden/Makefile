wasm: 
	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm ./wasm

server-app:
	go build -o server-app .

run: wasm server-app
	./server-app
