.PHONY: run build test clean

run:
	air

build:
	go build -o bin/go-kas

test:
	go test ./...

clean:
	rm -rf bin/
