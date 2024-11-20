#!/bin/zsh

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "🔄 Updating VS Code settings..."

if python3 "./scripts/maintain_vscode_settings.py"; then
    echo "${GREEN}✅ VS Code settings updated successfully!${NC}"

    # Optional: Show the diff
    git diff .vscode/settings.json

    # Optional: Automatically stage the changes
    read "REPLY?Do you want to stage the changes? (y/N) "
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        git add .vscode/settings.json
        echo "${GREEN}Changes staged!${NC}"
    fi
else
    echo "${RED}❌ Error updating VS Code settings${NC}"
    exit 1
fi
