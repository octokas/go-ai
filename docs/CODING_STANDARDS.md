# Go-Kas Coding Standards

## Code Organization

### Package Structure
```go
package example

// 1. Constants
const (
    DefaultTimeout = 30 * time.Second
    MaxRetries     = 3
)

// 2. Types
type Service struct {
    db     *database.DB
    logger *zap.Logger
    config *Config
}

// 3. Interfaces
type Repository interface {
    Find(ctx context.Context, id string) (*Model, error)
    Create(ctx context.Context, model *Model) error
}

// 4. Global Variables (avoid if possible)
var (
    errNotFound = errors.New("resource not found")
)
```

### Error Handling
```go
// 1. Custom Error Types
type Error struct {
    Code    string
    Message string
    Err     error
}

func (e *Error) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s: %v", e.Message, e.Err)
    }
    return e.Message
}

// 2. Error Creation
func NewNotFoundError(resource string) *Error {
    return &Error{
        Code:    "NOT_FOUND",
        Message: fmt.Sprintf("%s not found", resource),
    }
}

// 3. Error Handling Pattern
func (s *Service) GetTask(ctx context.Context, id string) (*Task, error) {
    task, err := s.repo.Find(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("finding task: %w", err)
    }
    return task, nil
}
```

### Context Usage
```go
// 1. Context with Values
type contextKey string

const (
    userIDKey contextKey = "userID"
    traceIDKey contextKey = "traceID"
)

// 2. Context Handling
func (s *Service) CreateTask(ctx context.Context, task *Task) error {
    // Always check context first
    if err := ctx.Err(); err != nil {
        return fmt.Errorf("context error: %w", err)
    }

    // Get values from context
    userID, ok := ctx.Value(userIDKey).(string)
    if !ok {
        return errors.New("user ID not found in context")
    }

    // Use context with timeout for external calls
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    return s.repo.Create(ctx, task)
}
```

### Logging Standards
```go
// 1. Logger Setup
func NewLogger() *zap.Logger {
    config := zap.NewProductionConfig()
    config.EncoderConfig.TimeKey = "timestamp"
    config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

    logger, _ := config.Build()
    return logger
}

// 2. Logging Pattern
func (s *Service) ProcessTask(ctx context.Context, taskID string) error {
    logger := s.logger.With(
        zap.String("taskID", taskID),
        zap.String("operation", "ProcessTask"),
    )

    logger.Info("starting task processing")

    if err := s.process(ctx, taskID); err != nil {
        logger.Error("failed to process task",
            zap.Error(err),
            zap.String("status", "failed"))
        return err
    }

    logger.Info("task processing completed",
        zap.String("status", "success"))
    return nil
}
```

### Testing Standards
```go
// 1. Table-Driven Tests
func TestService_ProcessTask(t *testing.T) {
    tests := []struct {
        name    string
        taskID  string
        setup   func(*testing.T) *Service
        wantErr bool
    }{
        {
            name:   "successful processing",
            taskID: "task-123",
            setup: func(t *testing.T) *Service {
                return NewTestService(t)
            },
            wantErr: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            svc := tt.setup(t)
            err := svc.ProcessTask(context.Background(), tt.taskID)

            if (err != nil) != tt.wantErr {
                t.Errorf("ProcessTask() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

## Design System Integration

### HTTP Response Format
```go
type Response struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   *ErrorInfo  `json:"error,omitempty"`
    Meta    *MetaInfo   `json:"meta,omitempty"`
}

type ErrorInfo struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
}

type MetaInfo struct {
    Page      int `json:"page,omitempty"`
    PageSize  int `json:"pageSize,omitempty"`
    TotalRows int `json:"totalRows,omitempty"`
}
```

### Design System Constants
```go
const (
    // Color scheme from our design system
    ColorPrimary   = "#0071e3"
    ColorSecondary = "#86868b"
    ColorError     = "#ff6188"
    ColorSuccess   = "#a9dc76"

    // Typography
    FontFamilySFPro = "SF Pro Text"
    FontSizeBase    = "15px"

    // Animation durations
    DurationFast    = 150
    DurationNormal  = 300
    DurationSlow    = 500
)
```
