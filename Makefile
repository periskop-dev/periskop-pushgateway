.PHONY: test lint

test:
	go test -v -race

lint :
	golangci-lint run

build:
	go build
