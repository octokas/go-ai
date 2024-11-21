.PHONY: all build run test clean docker

# Variables
BINARY_NAME=go-kas
DOCKER_IMAGE=go-kas

all: clean build

build:
	@echo "Building..."
	go build -o bin/$(BINARY_NAME) main.go

run:
	@echo "Running with air..."
	air

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning..."
	go clean
	rm -rf bin/

docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE)

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

lint:
	@echo "Linting..."
	golangci-lint run

dev: deps
	@echo "Starting development server..."
	air
