#!/bin/zsh

# Check for required tools
check_deps() {
    if ! command -v lolcat >/dev/null 2>&1; then
        echo "\033[38;5;209mInstalling lolcat for colorful output...\033[0m"
        brew install lolcat
    fi
    if ! command -v pv >/dev/null 2>&1; then
        echo "\033[38;5;209mInstalling pv for progress bars...\033[0m"
        brew install pv
    fi
}

check_deps

# Monokai Pro colors
# Purple: \033[38;5;141m
# Orange: \033[38;5;209m
# Green:  \033[38;5;108m
# Yellow: \033[38;5;222m
# Blue:   \033[38;5;73m

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

# Create debug directory in the project root
debug_dir="debug"
mkdir -p "$debug_dir"

# Fancy banner with Monokai Pro colors
echo "
\033[38;5;141m╔═══════════════════════════════════════╗
║\033[38;5;209m         Directory Map Generator       \033[38;5;141m║
║\033[38;5;108m         🗂  File Scanner 2024 🔍      \033[38;5;141m║
╚═══════════════════════════════════════╝\033[0m
" 

# Count total files for progress bar
echo "\033[38;5;222m🔍 Counting files...\033[0m"
total_files=$(find . -not \( -name ".git" -prune \) -not \( -name "debug" -prune \) -not \( -name "vendor" -prune \) -type f -o -type d | wc -l)

echo "\033[38;5;73m📂 Processing $total_files files and directories...\033[0m"

# Generate file data and pipe to jq with progress bar
find . -not \( -name ".git" -prune \) -not \( -name "debug" -prune \) -not \( -name "vendor" -prune \) -type f -o -type d | sort | \
pv -l -s $total_files -N "\033[38;5;141m🔍 Scanning  \033[0m" | \
while read -r file; do
    if [[ $file == "." ]]; then continue; fi
    
    name=${file##*/}
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
done | pv -l -s $total_files -N "\033[38;5;209m⚡️ Processing\033[0m" | jq -s '
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
' | pv -l -N "\033[38;5;108m💾 Writing   \033[0m" > "$debug_dir/directory_map.json"

echo "
\033[38;5;222m✨ Success! Directory map has been saved to:
\033[38;5;73m📁 $debug_dir/directory_map.json\033[0m
"