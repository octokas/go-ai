# Go Hello Dutonian

A simple starter web server in Go.

## Project Structure 

```
dutonian/
├── main.go
├── go.mod
├── HelloDutonian.md
└── deploy/
    ├── Dockerfile
    ├── docker-compose.yml
    └── .env.example
```

## Getting Started

### 1. Prerequisites

- `Go 1.x` or higher
- `Git`

### 2. Installation
```bash
# Clone the repository
git clone <your-repo-url>
cd dutonian

# Initialize the Go module
go mod init dutonian
```

### 3. Running the Server

```bash
# Run directly with Go
go run main.go

# Or use Air for hot-reloading during development
air
```

Visit `http://localhost:8080` in your browser to see "Hello, Dutonian! 👋"

### Development with Hot Reload

For hot reload functionality during development, install Air:

```bash
go install github.com/cosmtrek/air@latest
```

Then create a configuration file `.air.toml`:

```toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./tmp/main ."
  bin = "./tmp/main"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor"]
  include_ext = ["go"]
  exclude_regex = ["_test.go"]
```

Now you can run your server with hot reload:
```bash
air
```

## Deployment

The `deploy` directory contains everything needed for deploying the application:

### Local Docker Deployment

1. Copy the environment file:
```bash
cp deploy/.env.example deploy/.env
```

2. Build and run with Docker Compose:
```bash
docker-compose -f deploy/docker-compose.yml up --build
```

### Production Deployment

For production deployment:

1. Adjust environment variables in `.env` for production settings
2. Use Docker Compose:
```bash
docker-compose -f deploy/docker-compose.yml -f deploy/docker-compose.prod.yml up -d
```

The application will be available at `http://localhost:8080`

### Basic Deployment Options

You can also deploy this application in several other ways:

1. **Direct Server Deployment**
```bash
# Build the binary
go build -o server

# Run the server
./server
```

2. **Simple Docker Deployment**
```bash
# Build the Docker image
docker build -t hello-dutonian -f deploy/Dockerfile .

# Run the container
docker run -p 8080:8080 hello-dutonian
```

For production environments, remember to:
- Set appropriate environment variables
- Use HTTPS in production
- Consider using a process manager (like systemd) for automatic restarts

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details