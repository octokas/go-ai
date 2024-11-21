go-kas/
├── cmd/
│   └── server/
│       └── logging.go
│       └── server.go
│       └── ngrok.go
│       └── healthcheck.go
├── go/
│   ├── routers/
│   │   ├── routers.go      # Main router setup
│   │   ├── home.go         # Home route handlers
│   │   ├── tasks.go        # Task route handlers
│   │   └── assets.go       # Asset route handlers
│   │   └── users.go        # User route handlers
│   │   └── calendar.go    # Calendar route handlers
│   │   └── reports.go     # Report route handlers
│   │   └── apiv1.go      # API v1 routes
│   │   └── apiv2.go      # API v2 routes
│   │   └── graphql.go    # GraphQL routes
│   │   └── design.go     # Design routes
│   │   └── animation.go  # Animation routes
│   ├── internal/
│   │   ├── internal.go     # Internal package setup
│   │   ├── databases/
│   │   │   └── db.go       # Database connections
│   │   │   └── mongo.go    # MongoDB connections
│   │   │   └── sqlite.go   # SQLite connections
│   │   ├── apis/
│   │   │   └── api.go      # API routes
│   │   └── models/
│   │       ├── task.go
│   │       ├── subtask.go
│   │       └── asset.go
│   │       └── calendar.go
│   ├── handlers/
│   │   ├── handlers.go     # Common handler utilities
│   │   ├── tasks.go
│   │   └── assets.go
│   │   └── calendar.go
│   └── pkg/
│       ├── pkg.go          # Package utilities
│   │   └── config/
│   │       └── config.go   # Configuration utilities
│   │   └── design/
│   │       └── c4d.go      # C4D Animation utilities
│   │   └── animation/
│   │       └── animation.go   # Animation utilities
│   │       └── ae.go          # After Effects Animation utilities
│   │       └── calvary.go     # Calvary Animation utilities
│   │   └── design/
│   │       └── c4d.go         # C4D Design utilities
│   │       └── sketch.go     # Sketch Design utilities
│   │       └── sketch3D.go   # Sketch 3D Design utilities
│   │   └── database/
│   │       └── database.go   # Database utilities
│   │   └── reporting/
│   │       └── reporting.go   # Reporting utilities
│   │       └── security.go   # Security utilities
│   │       └── healthcheck.go   # Healthcheck utilities
│   │       └── accessibility.go   # Accessibility utilities
│   │       └── rate-limiting.go   # Rate limiting utilities
│   │   └── middleware/
│   │       └── middleware.go # Middleware utilities
│   ├── tests/
│   │   ├── tests.go
│   │   ├── server_tests.go
│   │   ├── router_tests.go
│   │   ├── handler_tests.go
│   │   ├── model_tests.go
│   │   ├── integration_tests.go
│   │   ├── unit_tests.go
├── design-system/          # Our design system
├── ui/                     # Hugo templates
└── Makefile
