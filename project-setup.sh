#!/bin/zsh

# Monokai Pro colors
PINK='\033[38;2;255;97;136m'
ORANGE='\033[38;2;255;159;67m'
YELLOW='\033[38;2;255;216;102m'
GREEN='\033[38;2;169;220;118m'
BLUE='\033[38;2;120;220;232m'
PURPLE='\033[38;2;187;128;255m'
WHITE='\033[38;2;252;252;252m'
NC='\033[0m'

# Function to enable color persistence like lolcat
setup_color_persistence() {
    # Add color persistence to shell RC file
    echo '# Monokai Pro color persistence
export FORCE_COLOR=true
export CLICOLOR=1
export CLICOLOR_FORCE=1

# Custom color formatting for common commands
alias ls="ls -G"
alias grep="grep --color=always"
alias less="less -R"
alias tree="tree -C"

# Force color in pipe
export LESS="-R"
export GREP_OPTIONS="--color=always"
export GREP_COLOR="1;31"

# Custom PS1 with Monokai Pro colors
export PS1="%F{#FF619B}%n%f@%F{#FFD866}%m%f:%F{#78DCE8}%~%f$ "' >> ~/.zshrc
}

# Install lolcat for additional color effects
command -v lolcat >/dev/null 2>&1 || { echo "Installing lolcat..."; brew install lolcat; }

echo "${BLUE}Setting up project structure...${NC}" | lolcat

# Check for required tools
echo "${YELLOW}Checking required tools...${NC}" | lolcat
command -v go >/dev/null 2>&1 || { echo "${ORANGE}Go is required but not installed. Installing...${NC}" | lolcat; brew install go; }
command -v air >/dev/null 2>&1 || { echo "${ORANGE}Air is required but not installed. Installing...${NC}" | lolcat; brew install cosmtrek/tools/air; }
command -v jq >/dev/null 2>&1 || { echo "${ORANGE}jq is required but not installed. Installing...${NC}" | lolcat; brew install jq; }
command -v yarn >/dev/null 2>&1 || { echo "${ORANGE}yarn is required but not installed. Installing...${NC}" | lolcat; brew install yarn; }

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
    main_file="${dir%/*}.go"
    touch "$dir/$main_file"
    echo "package ${dir%/*}" > "$dir/$main_file"
    echo "${GREEN}Created $dir and $dir/$main_file${NC}" | lolcat
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
    echo "package routers" > "routers/$file.go" | lolcat
done

# Create handler files
handler_files=("tasks" "assets" "users" "calendar" "reports")
for file in "${handler_files[@]}"; do
    touch "handlers/$file.go"
    echo "package handlers" > "handlers/$file.go" | lolcat
done

# Create model files
model_files=("task" "asset" "user" "calendar" "report")
for file in "${model_files[@]}"; do
    touch "models/$file.go"
    echo "package models" > "models/$file.go" | lolcat
done

# Create test files
test_files=("server" "router" "handler" "model" "integration" "unit")
for file in "${test_files[@]}"; do
    touch "tests/${file}_tests.go"
    echo "package tests" > "tests/${file}_tests.go" | lolcat
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
*.DS_Store
EOF

# Initialize git if not already initialized
if [ ! -d .git ]; then
    git init
    git add .
    git commit -m "Initial commit: Project structure setup"
fi


# Setup color persistence
setup_color_persistence

echo "${GREEN}Project structure setup complete!${NC}" | lolcat
echo "${BLUE}Next steps:${NC}" | lolcat
echo "${PINK}1. Review and customize the generated files${NC}" | lolcat
echo "${YELLOW}2. Run 'go mod tidy' to manage dependencies${NC}" | lolcat
echo "${GREEN}3. Start coding! Use 'make run' to start the development server${NC}" | lolcat
echo "${PURPLE}4. Color persistence has been added to your shell${NC}" | lolcat
echo "${ORANGE}5. Restart your terminal for color changes to take effect${NC}" | lolcat
