.PHONY: test
test:
	go test ./...

.PHONY: test
install:
	go install .

build:
	go build -o factor .
