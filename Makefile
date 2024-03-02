.DEFAULT_GOAL := pre-commit

gofmt:
	go fmt ./...

goimports:
	goimports -w .

fmt: goimports gofmt 
	
fix: fmt

lint:
	golangci-lint run ./...

test:
	go test -v ./...

bench:
	go test -bench=. ./...

ci: lint test bench

pre-commit: lint test

all: ci
