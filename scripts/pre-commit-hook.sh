#!/bin/zsh

# Path to your maintenance script
SCRIPT_PATH="./scripts/maintain_vscode_settings.py"

# Check if VS Code settings file has changed
if git diff --cached --name-only | grep -q ".vscode/settings.json"; then
    echo "📝 Formatting VS Code settings..."

    # Run the maintenance script
    if python3 "$SCRIPT_PATH"; then
        # Add the formatted file back to staging
        git add .vscode/settings.json
        echo "✅ VS Code settings formatted successfully!"
    else
        echo "❌ Error formatting VS Code settings"
        exit 1
    fi
fi
