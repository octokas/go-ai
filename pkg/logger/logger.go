package logger

import (
	"log"
	"net/http"
	"os"
	"sync"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct {
	logger *log.Logger
	level  Level
	mu     sync.Mutex
}

var defaultLogger *Logger
var once sync.Once

var New = func() *Logger {
	once.Do(func() {
		defaultLogger = &Logger{
			logger: log.New(os.Stdout, "", log.LstdFlags),
			level:  INFO,
		}
	})
	return defaultLogger
}

func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level <= DEBUG {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.logger.SetPrefix("[DEBUG] ")
		l.logger.Println(v...)
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.level <= INFO {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.logger.SetPrefix("[INFO] ")
		l.logger.Println(v...)
	}
}

func (l *Logger) Warn(v ...interface{}) {
	if l.level <= WARN {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.logger.SetPrefix("[WARN] ")
		l.logger.Println(v...)
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.level <= ERROR {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.logger.SetPrefix("[ERROR] ")
		l.logger.Println(v...)
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.level <= FATAL {
		l.mu.Lock()
		l.logger.SetPrefix("[FATAL] ")
		l.logger.Fatal(v...)
	}
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
