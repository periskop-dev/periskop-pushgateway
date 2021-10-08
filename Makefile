.PHONY: test lint build

test:
	go test -v -race

lint :
	golangci-lint run

build:
	go build
