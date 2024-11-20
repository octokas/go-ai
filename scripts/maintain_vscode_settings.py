#!/usr/bin/env python3
import json
import subprocess
from pathlib import Path

def get_installed_extensions():
    """Get list of installed VS Code extensions."""
    result = subprocess.run(['code', '--list-extensions'],
                          capture_output=True, text=True)
    return result.stdout.strip().split('\n')

def update_settings():
    """Update VS Code settings file with current extensions and formatting."""
    settings_path = Path('.vscode/settings.json')
    if not settings_path.exists():
        return False

    # Read current settings
    with open(settings_path, 'r') as f:
        settings = json.load(f)

    # Get current extensions
    extensions = get_installed_extensions()

    # Extension categories (maintain your current categorization)
    categories = {
        "Go Development": [ext for ext in extensions if any(x in ext.lower()
            for x in ['go', 'golang'])],
        "AWS Tools": [ext for ext in extensions if any(x in ext.lower()
            for x in ['aws', 'dynamodb', 'lambda'])],
        "C/C++/Objective-C": [ext for ext in extensions if any(x in ext.lower()
            for x in ['cpp', 'clang', 'lldb', 'c++'])],
        # ... add other categories
    }

    # Format settings with consistent spacing and ordering
    formatted_settings = format_settings(settings)

    # Write back formatted settings
    with open(settings_path, 'w') as f:
        json.dump(formatted_settings, f, indent=4)

    return True

def format_settings(settings):
    """Format settings with consistent spacing and categorization."""
    # Your custom formatting logic here
    return settings

if __name__ == "__main__":
    if update_settings():
        print("VS Code settings updated successfully!")
    else:
        print("Error updating VS Code settings")
        exit(1)
