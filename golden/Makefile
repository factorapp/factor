wasm: 
	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm .

run:  wasm
	go build -o server-app .
	./server-app
