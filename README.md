# 🤖 Go AI Project Template 

A modern template repository for building AI/ML applications in Go! 🚀

## ✨ Features

- Clean project structure optimized for AI/ML workloads
- Common AI dependencies pre-configured
- Docker support out of the box 🐳
- Example ML pipeline setup
- Testing framework ready to go ✅
- Automated changelog generation
- Security-first approach

## 🚀 Getting Started

1. Click "Use this template" to create your new repository
2. Clone your new repo locally
3. Run `go mod tidy` to install dependencies
4. Run `make init` to set up your development environment
5. Start with a simple example: `make run-hello`

### Quick Start Example
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Dutonian! Welcome to your AI project template!")
}
```

### Example ML Pipeline (coming soon)

## Project Structure
```
|-- cmd/ # Main application binary
|  |-- api/ # API definitions
|  |-- worker/ # Worker binary
|-- internal/ # Internal packages
|  |-- ai/ # AI specific logic
|  |-- pipeline/ # ML pipeline
|  |-- config/ # Configuration
|  |-- db/ # Database logic
|  |-- logger/ # Logging logic
|  |-- metrics/ # Metrics logic
|  |-- server/ # Server logic
|-- pkg/ # Reusable packages
|-- scripts/ # Build, maintenance, and utility scripts
|  |-- changelog.go # Changelog generator
|-- deploy/ # Deployment configuration
|  |-- Dockerfile
|  |-- docker-compose.yml
|  |-- .env.example
|-- docs/ # Documentation
|-- examples/ # Example code
|-- tests/ # Test suite
|-- .github/ # GitHub Actions configuration
|-- .gitignore # Git ignore file
|-- Dockerfile # Docker configuration
|-- go.mod # Go module file
|-- Makefile # Makefile for convenience
``` 

## 🛠️ Development

### Prerequisites

- Go 1.21 or higher
- Docker (optional)
- Make

### Common Commands

```bash
make init          ## Initialize development environment
make test          ## Run tests
make build         ## Build the project
make run-hello     ## Run the Hello Dutonian example
make changelog     ## Generate changelog
```

### Aggressive Code Cleanup

```bash
## More aggressive cache clearing
go clean -cache -modcache -i -r

## And/or delete the replace directive and run tidy again
go mod edit -dropreplace github.com/octokas/go-ai
go mod tidy       ## Clean the project
```

## 📝 Changelog

Changes are automatically tracked and documented in `CHANGELOG.md`. To generate the changelog:

```bash
go run scripts/changelog.go
```

The changelog generator scans commit messages following [Conventional Commits](https://www.conventionalcommits.org/) format.

## 🔒 Security

### Reporting Security Issues

We take security seriously. If you discover a security vulnerability, please follow these steps:

1. **Do NOT open a public issue**
2. Send a private email to security@yourdomain.com
3. Include detailed information about the vulnerability
4. Allow up to 48 hours for an initial response

### Security Best Practices

- All dependencies are automatically scanned for vulnerabilities
- Regular security audits are performed
- Updates are promptly released for security patches

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙋‍♂️ Getting Help

- Check out the [documentation](docs/README.md)
- Join our [Discord community](https://discord.gg/yourdiscord)
- Open an issue for bug reports or feature requests

