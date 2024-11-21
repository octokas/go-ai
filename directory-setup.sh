#!/bin/zsh

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo "${BLUE}Setting up go-kas project structure...${NC}"

# Check for required tools
echo "${BLUE}Checking required tools...${NC}"
command -v go >/dev/null 2>&1 || { echo "Go is required but not installed. Installing..."; brew install go; }
command -v air >/dev/null 2>&1 || { echo "Air is required but not installed. Installing..."; brew install cosmtrek/tools/air; }
command -v jq >/dev/null 2>&1 || { echo "jq is required but not installed. Installing..."; brew install jq; }
command -v yarn >/dev/null 2>&1 || { echo "yarn is required but not installed. Installing..."; brew install yarn; }

# Create main directories and their entry point files
directories=(
    "config"
    "logging"
    "server"
    "utils"
    "routers"
    "handlers"
    "services"
    "models"
    "tests"
    "databases/migrations"
    "reporting"
    "middleware"
    "apis"
    "pkg"
    "ae"
    "c4d"
    "sketch3d"
    "scripts"
)

# Create directories and their main Go files
for dir in "${directories[@]}"; do
    mkdir -p "$dir"
    main_file="${dir%/*}.go"  # Remove potential subdirectory for main file
    touch "$dir/$main_file"
    echo "package ${dir%/*}" > "$dir/$main_file"
    echo "${GREEN}Created $dir and $dir/$main_file${NC}"
done

# Create additional files in utils
touch utils/ngrok.go utils/healthcheck.go

# Create router files
router_files=(
    "home" "tasks" "assets" "users" "calendar" "reports"
    "apiv1" "apiv2" "graphql" "design" "animation"
)
for file in "${router_files[@]}"; do
    touch "routers/$file.go"
    echo "package routers" > "routers/$file.go"
done

# Create handler files
handler_files=("tasks" "assets" "users" "calendar" "reports")
for file in "${handler_files[@]}"; do
    touch "handlers/$file.go"
    echo "package handlers" > "handlers/$file.go"
done

# Create model files
model_files=("task" "asset" "user" "calendar" "report")
for file in "${model_files[@]}"; do
    touch "models/$file.go"
    echo "package models" > "models/$file.go"
done

# Create test files
test_files=("server" "router" "handler" "model" "integration" "unit")
for file in "${test_files[@]}"; do
    touch "tests/${file}_tests.go"
    echo "package tests" > "tests/${file}_tests.go"
done

# Create database files
touch databases/mongo.go databases/sqlite.go databases/migrations.go

# Create reporting files
touch reporting/security.go reporting/healthcheck.go reporting/accessibility.go

# Create middleware files
touch middleware/rate-limiting.go

# Create API files
touch apis/v1.go apis/v2.go apis/graphql.go

# Create root level files
touch main.go go.mod go.sum Makefile

# Initialize go.mod
go mod init go-kas

# Create basic Makefile
cat > Makefile << 'EOF'
.PHONY: run build test clean

run:
	air

build:
	go build -o bin/go-kas

test:
	go test ./...

clean:
	rm -rf bin/
EOF

# Create basic .gitignore
cat > .gitignore << 'EOF'
# Binaries
bin/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# IDE specific files
.idea/
.vscode/
*.swp
*.swo

# Air tmp directory
tmp/

# OS specific files
.DS_Store
EOF

# Initialize git if not already initialized
if [ ! -d .git ]; then
    git init
    git add .
    git commit -m "Initial commit: Project structure setup"
fi

echo "${GREEN}Project structure setup complete!${NC}"
echo "${BLUE}Next steps:${NC}"
echo "1. Review and customize the generated files"
echo "2. Run 'go mod tidy' to manage dependencies"
echo "3. Start coding! Use 'make run' to start the development server"
