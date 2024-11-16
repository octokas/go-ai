# API Router Documentation 🛣️

## Overview 📝
Our API uses a modular routing system organized by feature and version. All routes are defined in the `pkg/router` directory.

## Structure 🏗️
```
pkg/router/
├── router.go    # Main router setup
├── home.go      # Home/static page routes
├── api_v1.go    # Version 1 API routes
└── api_v2.go    # Version 2 API routes
```

## Current Routes 🚏

### Home Routes 🏠
- `/` - Welcome page
- `/about` - About us information
- `/contact` - Contact details

### API v1 Routes 🔵
- `/api/v1/hello` - Basic greeting
- `/api/v1/users` - List users (basic info)
- `/api/v1/status` - Basic system status

### API v2 Routes 🟢
- `/api/v2/hello` - Enhanced greeting with version
- `/api/v2/users` - List users (detailed info)
- `/api/v2/status` - Detailed system status with services

## Adding New Routes 🆕

### Adding a Static Page
```go
// In pkg/router/home.go
func setupHomeRoutes() {
    // ... existing routes ...
    http.HandleFunc("/new-page", newPageHandler)
}

func newPageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the new page! ✨")
}
```

### Adding a New API Endpoint
```go
// In pkg/router/api_v2.go
func setupV2Routes() {
    // ... existing routes ...
    http.HandleFunc("/api/v2/products", v2ProductsHandler)
}

func v2ProductsHandler(w http.ResponseWriter, r *http.Request) {
    products := []map[string]interface{}{
        {
            "id": "1",
            "name": "Awesome Product",
            "price": 99.99,
            "inStock": true,
        },
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}
```

### Adding a New API Version 🔄
1. Create a new file `api_v3.go`:
```go
package router

import (
    "encoding/json"
    "net/http"
)

func setupV3Routes() {
    http.HandleFunc("/api/v3/hello", v3HelloHandler)
    // Add more v3 routes...
}

func v3HelloHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]interface{}{
        "message": "Hello from API v3! 👋",
        "version": 3,
        "features": []string{"feature1", "feature2"},
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
```

2. Update `router.go`:
```go
func Setup() {
    setupHomeRoutes()
    setupV1Routes()
    setupV2Routes()
    setupV3Routes() // Add the new version
}
```

## Best Practices 🌟

1. **Version Consistency** 📊
   - Keep response formats consistent within each API version
   - Document breaking changes between versions

2. **Response Headers** 📫
   - Always set appropriate Content-Type headers
   - Consider adding API version headers

3. **Error Handling** ⚠️
```go
func exampleHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    if err := someOperation(); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{
            "error": err.Error(),
        })
        return
    }
    
    // Success response...
}
```

4. **Route Naming** 📝
   - Use clear, descriptive route names
   - Follow RESTful conventions where appropriate
   - Keep versioning consistent


## Testing Routes 🧪

### Automated Tests
We use Go's testing package to ensure our routes work correctly. Here's an example of how to test routes:
```go
func TestHomeHandler(t *testing.T) {
	// create a request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)

	// our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder
	handler.ServeHTTP(rr, req)

	// check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check the response body is what we expect
	expected := "Welcome to the Home Page! 🏠"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
```

_Run the tests using:_
```bash
go test ./pkg/router -v
```

_You can test routes using curl:_
```bash
## Test v1 endpoint
curl http://localhost:8080/api/v1/hello

## Test v2 endpoint
curl http://localhost:8080/api/v2/users

## Test v2 endpoint with pretty-printed JSON
curl http://localhost:8080/api/v2/users | jq '.'
```

_Or using the browser for GET requests:_
- http://localhost:8080/
- http://localhost:8080/api/v1/status
- http://localhost:8080/api/v2/status


## Future Considerations 🔮

1. **Middleware Support** 🔄
   - Authentication
   - Logging
   - Rate limiting

2. **API Documentation** 📚
   - Consider adding Swagger/OpenAPI specs
   - Automated documentation generation

3. **Versioning Strategy** 📈
   - URL versioning (current)
   - Header versioning
   - Content negotiation

4. **Performance Optimization** ⚡
   - Route caching
   - Response compression
   - Connection pooling

## Need Help? 💡
Contact the development team or create an issue in the repository for:
- Bug reports
- Feature requests
- Documentation improvements