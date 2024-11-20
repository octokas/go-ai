## Setting Up for VSCode

You can list your VS Code extensions in two ways:

1. First, export your current extensions list by running this command in your terminal:
```bash
code --list-extensions > .vscode/extensions.list
```

2. Then, create an `extensions.json` file in the `.vscode` directory with the following content:
```json
{
    "recommendations": [
        "dbaeumer.vscode-eslint",
        "esbenp.prettier-vscode",
        "golang.go",
        "mongodb.mongodb-vscode",
        "ms-vscode.cpptools"
    ]
}
```

3. Finally, you can automatically install the extensions by running this command in your terminal:
```bash
./scripts/install-vscode-extensions.sh
```

When anyone opens this project in VS Code, they will have the extensions you listed in the `extensions.json` file automatically installed.

To get your current extensions in the correct format, you can run the following command in your terminal:
```bash
## using sed to format the output
code --list-extensions | sed 's/^/"&",/' > .vscode/extensions.list

## using awk to format the output
code --list-extensions | awk -v q="'" '{printf "%s%s%s\n", q, $0, q}' > .vscode/extensions.list

## using jq to format the output
code --list-extensions | jq -R -s -c 'split("\n")[:-1]' > .vscode/extensions.list
```
