go-kas/
├── config/ # for configuration logic
│   └── config.go ## (serves as main for config)
├── logging/ # for logging concerns
│   └── logging.go ## (serves as entry point for logging)
├── server/ # for server logic
│   └── server.go ## (serves as entry point for server)
├── utils/ # for utility functions
│   └── utils.go ## (serves as entry point for utils)
│   └── ngrok.go
│   └── healthcheck.go
├── routers/ # for handling routing logic
│   ├── routers.go ## (serves as entry point for routers)
│   └── home.go
│   └── tasks.go
│   └── assets.go
│   └── users.go
│   └── calendar.go
│   └── reports.go
│   └── apiv1.go
│   └── apiv2.go
│   └── graphql.go
│   └── design.go
│   └── animation.go
├── handlers/ # for request and response logic
│   └── handlers.go ## (serves as entry point for handlers)
│   └── tasks.go
│   └── assets.go
│   └── users.go
│   └── calendar.go
│   └── reports.go
├── services/ # for business logic
│   └── services.go ## (serves as entry point for services)
├── models/ # for data structures/models
│   └── models.go ## (serves as entry point for models)
│   └── task.go
│   └── asset.go
│   └── user.go
│   └── calendar.go
│   └── report.go
├── tests/ # for comprehensive testing logic
│   ├── tests.go ## (serves as entry point for tests)
│   ├── server_tests.go
│   ├── router_tests.go
│   ├── handler_tests.go
│   ├── model_tests.go
│   ├── integration_tests.go
│   ├── unit_tests.go
├── databases/ # for database logic with migrations support
│   ├── databases.go ## (serves as entry point for databases)
│   ├── mongo.go
│   ├── sqlite.go
│   ├── migrations/ # for database schema migrations
│   ├── migrations.go ## (serves as entry point for migrations)
├── reporting/ # for reporting concerns
│   ├── reporting.go ## (serves as entry point for reporting)
│   ├── security.go
│   ├── healthcheck.go
│   ├── accessibility.go
├── middleware/ # for cross-cutting concerns
│   ├── middleware.go ## (serves as entry point for middleware)
│   ├── rate-limiting.go
├── apis/ # for api logic
│   ├── api.go ## (serves as entry point for apis)
│   ├── v1.go
│   ├── v2.go
│   ├── graphql.go
├── pkg/ # for shared packages
│   ├── pkg.go ## (serves as entry point for packages)
├── ae/ # for after effects
│   ├── ae.go ## (serves as entry point for after effects)
├── c4d/ # for c4d
│   ├── c4d.go ## (serves as entry point for c4d)
├── sketch3d/ # for sketch3d
│   ├── sketch3d.go ## (serves as entry point for sketch3d)
├── design-system/ # our design system
├── docs/ # guides and documentation
├── hugo/ # hugo frontend ui
├── assets/ # static assets
├── scripts/ # scripts for automation
├── main.go # essential entry point
├── README.md # project documentation
├── go.mod # module declaration
├── go.sum # dependency resolution
└── Makefile # automation script for building and running the project
