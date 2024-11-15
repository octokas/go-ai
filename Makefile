.PHONY: all build test clean run help

# Variables
BINARY_NAME=myapp
BUILD_DIR=build
MAIN_API=cmd/api/main.go
MAIN_WORKER=cmd/worker/main.go

# Default target
all: build

# Build the project
build:
	@echo "Building API..."
	go build -o $(BUILD_DIR)/api $(MAIN_API)
	@echo "Building Worker..."
	go build -o $(BUILD_DIR)/worker $(MAIN_WORKER)

# Run tests
test:
	go test ./... -v -cover

# Clean build directory
clean:
	rm -rf $(BUILD_DIR)

# Run the API
run-api:
	go run $(MAIN_API)

# Run the worker
run-worker:
	go run $(MAIN_WORKER)

# Generate changelog
changelog:
	go run scripts/changelog.go

# Initialize development environment
init:
	go mod tidy
	go mod verify

# Initialize GitHub repository
init-repo:
	@echo "Initializing GitHub repository..."
	git init
	git checkout -b trunk 2>/dev/null || git checkout trunk
	@if gh repo create octokas/go-ai --private --source . --remote origin 2>/dev/null; then \
		echo "Repository created successfully"; \
	else \
		echo "Repository already exists, setting remote origin..."; \
		git remote add origin git@github.com:octokas/go-ai.git || true; \
	fi
	git add .
	git commit -m "Initial commit"
	git push -u origin trunk
	@echo "Repository initialized at github.com/octokas/go-ai"

# Help command
help:
	@echo "Available commands:"
	@echo "  make build      - Build the project"
	@echo "  make test       - Run tests"
	@echo "  make clean      - Clean build directory"
	@echo "  make run-api    - Run the API"
	@echo "  make run-worker - Run the worker"
	@echo "  make changelog  - Generate changelog"
	@echo "  make init       - Initialize development environment"
	@echo "  make init-repo  - Initialize GitHub repository"

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out 