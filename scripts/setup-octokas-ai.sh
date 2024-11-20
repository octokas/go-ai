#!/bin/zsh

##########################################
# Setup Octokas-AI #######################
##########################################

# Set strict error handling ##############
set -e

# Configuration ##########################
APP_NAME="octokas-ai"
GITHUB_WORKFLOW_DIR=".github/workflows"
CHANGES_DIR="changes"
SCRIPTS_DIR="scripts"
CURRENT_DATE=$(date '+%Y-%m-%d')
VERSION="0.1.0"

# Create main project structure ###########
create_project_structure() {
    echo "Creating project structure for $APP_NAME..."

    # Create main directories ###############
    mkdir -p $APP_NAME/{cmd,internal,pkg,api,configs,scripts,changes}
    mkdir -p $APP_NAME/$GITHUB_WORKFLOW_DIR

    # Create internal structure ############
    mkdir -p $APP_NAME/internal/{models,services,middleware,database}

    # Create main.go #######################
    cat > $APP_NAME/cmd/main.go << 'EOL'
package main

import (
    "log"
)

func main() {
    log.Println("Starting Octokas-AI server...")
}
EOL

    # Create .gitignore #####################
    cat > $APP_NAME/.gitignore << 'EOL'
# Environment variables
.env
.env.*
!.env.example

# Binary files
bin/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with go test -c
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Air temporary files
tmp/

# MongoDB data directory
data/db/

# IDE specific files
.idea/
.vscode/
*.swp
*.swo
EOL

    # Create .env.example ###################
    cat > $APP_NAME/.env.example << 'EOL'
# Server Configuration
PORT=8080
ENV=development

# MongoDB Configuration
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=octokas_ai

# Application Secrets
APP_SECRET=your_secret_here
EOL

    # Create actual .env file #################
    cp $APP_NAME/.env.example $APP_NAME/.env
}

# Create GitHub Actions workflow ##############
create_github_workflow() {
    cat > $APP_NAME/$GITHUB_WORKFLOW_DIR/main.yml << 'EOL'
name: CI/CD Pipeline

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'

    - name: Build
      run: go build -v ./...

    - name: Test
      run: go test -v ./...

    - name: Store Secrets
      run: |
        echo "${{ secrets.ENV_FILE }}" > .env
EOL
}

# Create MongoDB schema for file tracking #####
create_mongo_schema() {
    cat > $APP_NAME/internal/database/schema.go << 'EOL'
package database

type FileMetadata struct {
    Path      string   `bson:"path"`
    Name      string   `bson:"name"`
    Type      string   `bson:"type"`
    Size      int64    `bson:"size"`
    UpdatedAt string   `bson:"updated_at"`
    DataTypes []string `bson:"data_types"`
}
EOL
}

# Create dependency management script #####
create_dependency_script() {
    cat > $APP_NAME/$SCRIPTS_DIR/setup_deps.sh << 'EOL'
#!/bin/zsh

# Check for Homebrew installation
if ! command -v brew &> /dev/null; then
    echo "Installing Homebrew..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

# Install jq if not present
if ! command -v jq &> /dev/null; then
    echo "Installing jq..."
    brew install jq
fi

# Install Go if not present
if ! command -v go &> /dev/null; then
    echo "Installing Go..."
    brew install go
fi

# Install Air for live reload
if ! command -v air &> /dev/null; then
    echo "Installing Air..."
    /bin/bash -c "$(curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin)"
    export PATH=$PATH:$(go env GOPATH)/bin
    echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.zshrc
    air init
fi

# Install MongoDB
if ! command -v mongod &> /dev/null; then
    echo "Installing MongoDB..."
    brew tap mongodb/brew
    brew install mongodb-community
fi

# Install project dependencies
go mod tidy
EOL
    chmod +x $APP_NAME/$SCRIPTS_DIR/setup_deps.sh
}

# Create pre-commit script #################
create_pre_commit_script() {
    cat > $APP_NAME/$SCRIPTS_DIR/pre-commit.sh << 'EOL'
#!/bin/zsh

# Run tests
go test ./...

# Format code
go fmt ./...

# Run linter
if command -v golangci-lint &> /dev/null; then
    golangci-lint run
fi

# Update changelog
./scripts/update_changelog.sh
EOL
    chmod +x $APP_NAME/$SCRIPTS_DIR/pre-commit.sh
}

# Create cache clearing script #################
create_cache_clear_script() {
    cat > $APP_NAME/$SCRIPTS_DIR/clear_cache.sh << 'EOL'
#!/bin/zsh

# Clear Go cache
go clean -cache -modcache -i -r

# Clear Air tmp directory
if [ -d "tmp" ]; then
    rm -rf tmp
fi

# Clear MongoDB data (optional)
echo "Do you want to clear MongoDB data? (y/n)"
read -r response
if [[ $response =~ ^([yY][eE][sS]|[yY])$ ]]; then
    brew services stop mongodb-community
    rm -rf /usr/local/var/mongodb
    mkdir -p /usr/local/var/mongodb
    brew services start mongodb-community
fi
EOL
    chmod +x $APP_NAME/$SCRIPTS_DIR/clear_cache.sh
}

# Create changelog script ##################
create_changelog_script() {
    cat > $APP_NAME/$SCRIPTS_DIR/update_changelog.sh << 'EOL'
#!/bin/zsh

CHANGES_DIR="changes"
DATE=$(date '+%Y-%m-%d')
CHANGELOG_FILE="$CHANGES_DIR/$DATE.md"

# Create changes directory if it doesn't exist
mkdir -p $CHANGES_DIR

# Get the last commit message and diff
LAST_COMMIT_MSG=$(git log -1 --pretty=%B)
DIFF_SUMMARY=$(git diff HEAD^..HEAD --stat)

# Create or append to changelog
cat >> $CHANGELOG_FILE << EOF

## $(date '+%Y-%m-%d %H:%M:%S')

### Commit Message
$LAST_COMMIT_MSG

### Changes
\`\`\`
$DIFF_SUMMARY
\`\`\`
EOF
EOL
    chmod +x $APP_NAME/$SCRIPTS_DIR/update_changelog.sh
}

# Create air configuration ##################
create_air_config() {
    cat > $APP_NAME/.air.toml << 'EOL'
root = "."
tmp_dir = "tmp"

[build]
cmd = "go build -o ./tmp/main ./cmd/main.go"
bin = "tmp/main"
full_bin = "APP_ENV=dev ./tmp/main"
include_ext = ["go", "tpl", "tmpl", "html"]
exclude_dir = ["assets", "tmp", "vendor"]
include_dir = []
exclude_file = []
delay = 1000
stop_on_error = true
log = "air.log"

[log]
time = false

[color]
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"
EOL
}

# Initialize Git repository ##################
init_git_repo() {
    cd $APP_NAME
    git init
    git add .
    git commit -m "Initial commit: Project structure setup"
}

# Main execution ############################
main() {
    create_project_structure
    create_github_workflow
    create_mongo_schema
    create_dependency_script
    create_pre_commit_script
    create_cache_clear_script
    create_changelog_script
    create_air_config
    init_git_repo

    echo "Project $APP_NAME has been successfully created!"
    echo "To get started:"
    echo "1. cd $APP_NAME"
    echo "2. ./scripts/setup_deps.sh"
    echo "3. air"
}

main


##########################################
# 💞 @octokas ############################
##########################################
