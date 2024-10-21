.PHONY: clean lint test build

default: clean lint test build

clean:
	rm -rf dist/ cover.out

test: clean
	go test -v -cover ./...

lint:
	golangci-lint run

build:
	go build -ldflags "-s -w" -trimpath ./cmd/filen/
