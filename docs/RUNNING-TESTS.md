# Go Testing Guide

## Running Tests

```bash
# Run all tests in current directory and subdirectories
go test ./...

# Run all tests with verbose output (-v)
go test -v ./...

# Run all tests and show coverage
go test -cover ./...

# Run all tests with verbose output and coverage
go test -v -cover ./...

# Run specific test function
go test -run TestFunctionName

# Run tests with race condition detection
go test -race ./...

# Generate coverage report in HTML format
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## Writing Effective Tests

### Test File Naming
- Test files should end with `_test.go`
- Place test files in the same directory as the code they test
- Example: `user.go` → `user_test.go`

### Test Function Naming
```go
// Test functions should begin with "Test"
func TestCreateUser(t *testing.T) { ... }

// Table-driven tests are recommended for multiple test cases
func TestCalculate(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"positive", 1, 2},
        {"zero", 0, 1},
        {"negative", -1, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Calculate(tt.input)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}
```

### Best Practices
1. Use `t.Parallel()` for concurrent test execution when possible
2. Use subtests with `t.Run()` for better organization
3. Use test helpers for common setup/teardown
4. Use meaningful test names that describe the scenario
5. Include both positive and negative test cases
6. Use `testdata` directory for test fixtures
7. Use `assert` package from `testify` for cleaner assertions

### Test Structure
```go
func TestSomething(t *testing.T) {
    // Setup
    // ... prepare test data

    // Teardown (if needed)
    defer func() {
        // ... cleanup
    }()

    // Execute
    result := FunctionUnderTest()

    // Assert
    if expected != result {
        t.Errorf("expected %v, got %v", expected, result)
    }
}
```
