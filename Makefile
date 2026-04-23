.PHONY: all build run clean fmt lint test

all: lint fmt build

build:
	go build -o splanes .

run:
	go run .

clean:
	rm -f splanes

fmt:
	gofumpt -w .

lint:
	golangci-lint run ./...

test:
	go test ./...