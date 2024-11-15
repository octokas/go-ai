.PHONY: all build test clean run help dev init-repo init tidy
dev:
	air

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
	go clean -modcache

# Tidy up the project
tidy:
	go mod tidy

# Run with hot reload
dev:
	air

# Run the API
run-api:
	go run $(MAIN_API)

# Run the worker
run-worker:
	go run $(MAIN_WORKER)

# Generate and commit changelog
changelog:
	@echo "Generating changelog..."
	@mkdir -p changelogs
	@COMMIT_ID=$$(git rev-parse HEAD) && \
	AUTHOR=$$(git config user.name) && \
	VERSION=$$(date +%Y%m%d_%H%M%S)_$${COMMIT_ID:0:8} && \
	go run scripts/changelog.go > changelogs/changelog_$${VERSION}.md && \
	git add changelogs/changelog_$${VERSION}.md && \
	git commit -m "docs: changelog for version $${VERSION} by $${AUTHOR}"

# Initialize development environment
init:
	go mod tidy
	go mod verify

# Initialize GitHub repository with pre-push hook
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
	@echo "#!/bin/sh\nmake changelog" > .git/hooks/pre-push
	@chmod +x .git/hooks/pre-push
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