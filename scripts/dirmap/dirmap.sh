#!/bin/zsh

# Check for required tools
check_deps() {
    if ! command -v lolcat >/dev/null 2>&1; then
        echo "Installing lolcat for colorful output..."
        brew install lolcat
    fi
    if ! command -v pv >/dev/null 2>&1; then
        echo "Installing pv for progress bars..."
        brew install pv
    fi
}

check_deps

# Function to detect file language based on extension
get_language() {
    case ${1:l} in
        *.go)    echo "Go" ;;
        *.js)    echo "JavaScript" ;;
        *.ts)    echo "TypeScript" ;;
        *.jsx)   echo "React/JavaScript" ;;
        *.tsx)   echo "React/TypeScript" ;;
        *.py)    echo "Python" ;;
        *.rb)    echo "Ruby" ;;
        *.java)  echo "Java" ;;
        *.html)  echo "HTML" ;;
        *.css)   echo "CSS" ;;
        *.scss)  echo "SCSS" ;;
        *.sql)   echo "SQL" ;;
        *.md)    echo "Markdown" ;;
        *.json)  echo "JSON" ;;
        *.yaml|*.yml) echo "YAML" ;;
        *.sh|*.bash)  echo "Shell" ;;
        *.php)   echo "PHP" ;;
        *.rs)    echo "Rust" ;;
        *.cpp)   echo "C++" ;;
        *.c)     echo "C" ;;
        *.swift) echo "Swift" ;;
        *.kt)    echo "Kotlin" ;;
        *)       echo "null" ;;
    esac
}

# Create debug directory
root_dir=$(pwd)
debug_dir="${root_dir}/debug"
mkdir -p "$debug_dir"

# Fancy banner
echo "
╔═══════════════════════════════════════╗
║         Directory Map Generator       ║
║         🗂  File Scanner 2024 🔍      ║
╚═══════════════════════════════════════╝
" | lolcat -a -d 1

# Count total files for progress bar
echo "🔍 Counting files..." | lolcat -a
total_files=$(find . -not \( -name ".git" -prune \) -not \( -name "debug" -prune \) -type f -o -type d | wc -l)

echo "📂 Processing $total_files files and directories..." | lolcat -a

# Generate file data and pipe to jq with progress bar
find . -not \( -name ".git" -prune \) -not \( -name "debug" -prune \) -type f -o -type d | sort | \
pv -l -s $total_files -N "🔍 Scanning  " | \
while read -r file; do
    if [[ $file == "." ]]; then continue; fi
    
    name=$(basename "$file")
    path=${file#./}
    
    if [[ -d $file ]]; then
        type="directory"
        size=0
        echo "{\"name\":\"$name\",\"path\":\"$path\",\"type\":\"$type\",\"size\":$size}"
    else
        type="file"
        size=$(stat -f %z "$file" 2>/dev/null || stat -c %s "$file" 2>/dev/null)
        lang=$(get_language "$name")
        if [[ $lang == "null" ]]; then
            echo "{\"name\":\"$name\",\"path\":\"$path\",\"type\":\"$type\",\"size\":$size,\"language\":null}"
        else
            echo "{\"name\":\"$name\",\"path\":\"$path\",\"type\":\"$type\",\"size\":$size,\"language\":\"$lang\"}"
        fi
    fi
done | pv -l -s $total_files -N "⚡️ Processing" | jq -s '
    def nest($items):
        reduce ($items[] | select(.path != null) | {
            key: .path | split("/"),
            value: {
                name: .name,
                path: .path,
                type: .type,
                size: .size,
                language: .language
            }
        }) as $item ({};
            setpath($item.key; $item.value)
        );
    nest(.)
' | pv -l -N "💾 Writing   " > "$debug_dir/directory_map.json"

echo "
✨ Success! Directory map has been saved to:
📁 $debug_dir/directory_map.json
" | lolcat -a -d 1 