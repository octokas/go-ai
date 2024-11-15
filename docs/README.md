# Go Worker Template

A simple, configurable worker template in Go.


## Project Structure 

```
go-ai/
├── cmd/
│   └── api/
│       └── main.go
│   └── worker/
│       └── main.go
├── deploy/
│   └── Dockerfile
├── internal/
│   └── ai/
│       └── model.go
│       └── model_test.go
│   └── config/
│       └── config.go
│       └── config_test.go
│   └── logger/
│       └── logger.go
│       └── logger_test.go
│   └── metrics/
│       └── metrics.go
│       └── metrics_test.go
│   └── pipeline/
│       └── pipeline.go
│       └── pipeline_test.go
│   └── server/
│       └── server.go
│       └── server_test.go
│   └── db/
│       └── db.go
│       └── db_test.go
├── go.mod
├── go.sum
├── .gitignore
├── .env
├── .gitconfig
├── docs/
│   └── README.md
├── .github/
│   └── workflows/
│       └── ci.yml
```

## Getting Started

### 1. Prerequisites

- `Go 1.x` or higher
- `Git`

### 2. Installation
```bash
## Clone the repository
git clone https://github.com/octokas/go-ai
cd go-ai
## Install dependencies
go mod download
## Ensure you have a `.env` file in the root of the project
go env -w GITHUB_TOKEN=<your-github-token>
## Set the GITHUB_TOKEN environment variable
export GITHUB_TOKEN=<your-github-token>
## Set the GITHUB_TOKEN environment variable in the `.env` file
echo "GITHUB_TOKEN=<your-github-token>" >> .env
## Set the GITHUB_TOKEN environment variable in the `.gitconfig` file
echo "url \"https://${GITHUB_TOKEN}@github.com/\"" >> .gitconfig
## Ensure module dependencies are installed
go mod tidy
```

### 3. Configuration

Create a configuration file (e.g., `config.yaml` or use environment variables based on your `config` implementation):

```yaml
## Example config.yaml
worker:
  name: "my-worker"
  interval: "5s"
```

### 4. Running the Worker

```bash
# Run directly with Go
go run cmd/worker/main.go

# Or build and run the binary
go build -o worker cmd/worker/main.go
./worker
```

## Customizing the Worker

### 1. Add Configuration Options

Edit `internal/config/config.go` to add your configuration options:

```go
type Config struct {
    WorkerName string `yaml:"worker_name"`
    // Add your configuration fields here
}
```

### 2. Implement Worker Logic

Edit `cmd/worker/main.go` and add your processing logic in the `Run()` method:

```go
func (w *Worker) Run() error {
    w.logger.Info("Worker pipeline started")
    
    // Add your worker logic here
    // Example:
    // - Process files
    // - Make API calls
    // - Handle database operations
    
    return nil
}
```

### 3. Error Handling

The template includes basic error handling:
- Configuration errors will stop the worker
- Runtime errors in `Run()` will be logged and stop the worker

## Best Practices

1. Keep the `main.go` file focused on initialization and coordination
2. Add business logic in separate packages under `internal/`
3. Use the logger for meaningful operational insights
4. Handle graceful shutdowns when implementing long-running processes

## Example Implementation

Here's a minimal example adding a processing function:

```go
func (w *Worker) processItem() error {
    w.logger.Info("Processing item...")
    // Your processing logic here
    return nil
}

func (w *Worker) Run() error {
    w.logger.Info("Worker pipeline started")
    
    if err := w.processItem(); err != nil {
        return fmt.Errorf("failed to process item: %w", err)
    }
    
    return nil
}
```

## Helpful Commands

```bash
## Ensure you have a `.env` file in the root of the project
go env -w GITHUB_TOKEN=<your-github-token>

## Set the GITHUB_TOKEN environment variable
export GITHUB_TOKEN=<your-github-token>

## Set the GITHUB_TOKEN environment variable in the `.env` file
echo "GITHUB_TOKEN=<your-github-token>" >> .env

## Set the GITHUB_TOKEN environment variable in the `.gitconfig` file
echo "url \"https://${GITHUB_TOKEN}@github.com/\"" >> .gitconfig

## Ensure module dependencies are installed
go mod tidy

## Git pull the latest changes
git pull

## Git Pre-commit Hooks
pre-commit install

## Go Cleanup
go clean

## Run the tests
go test ./...

## Build the project
go build -o worker cmd/worker/main.go

## Run the worker
./worker

## Run the API
go run cmd/api/main.go

## Run the worker and API
./worker & go run cmd/api/main.go

## Run the worker, API, and tests
./worker & go run cmd/api/main.go && go test ./...

## Run the worker, API, and tests in watch mode
./worker & go run cmd/api/main.go && go test -watch ./...

## Run the worker, API, and tests in watch mode with verbose output
./worker & go run cmd/api/main.go && go test -v -watch ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details
