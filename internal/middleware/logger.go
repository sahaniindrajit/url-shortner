package middleware

import (
	"log"
	"net/http"
	"time"
)

// responseRecorder wraps http.ResponseWriter to capture status code
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

// Logger middleware logs request method, path, status and duration
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // default
		}

		next.ServeHTTP(recorder, r)

		duration := time.Since(start)

		log.Printf(
			"%s %s → %d (%s)",
			r.Method,
			r.URL.Path,
			recorder.statusCode,
			duration,
		)
	})
}
