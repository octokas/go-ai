#!/usr/bin/env python3
import re
import sys
from pathlib import Path

def fix_yaml_comments(file_path):
    # Skip markdown files
    if file_path.suffix.lower() in ['.md', '.markdown']:
        return False
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    # Only process files that look like YAML
    if not any(content.startswith(prefix) for prefix in ['---', 'name:', 'on:']):
        return False

    # Replace single # comments with ## but ignore if already ##
    modified_content = re.sub(r'(?m)^([^#]*)(?<![#])#(?!#)\s', r'\1## ', content)
    
    if modified_content != content:
        with open(file_path, 'w') as f:
            f.write(modified_content)
        return True
    
    return False

def main():
    changed = False
    for file_path in sys.argv[1:]:
        if fix_yaml_comments(Path(file_path)):
            print(f"Fixed comments in {file_path}")
            changed = True
    
    sys.exit(1 if changed else 0)

if __name__ == '__main__':
    main() 