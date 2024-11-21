package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

type Middleware func(http.Handler) http.Handler

func MiddlewareChain(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

func MiddlewareLogger(logger *log.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Create a custom response writer to capture status code
			rw := &ResponseWriter{w, http.StatusOK}

			next.ServeHTTP(rw, r)

			logger.Printf(
				"%s %s %s %d %v",
				r.Method,
				r.RequestURI,
				r.RemoteAddr,
				rw.status,
				time.Since(start),
			)
		})
	}
}

func MiddlewareRecovery(logger *log.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Printf("panic: %v\n%s", err, debug.Stack())
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

type ResponseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
