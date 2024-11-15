#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
    export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
fi

# Check if GITHUB_TOKEN is set
if [ -z "$GITHUB_TOKEN" ]; then
    echo "Error: GITHUB_TOKEN is not set in .env file"
    exit 1
fi

# Configure git to use the token
git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

echo "Git configured successfully with GitHub token"

# Run go mod tidy
go mod tidy 