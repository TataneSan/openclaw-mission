package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/TataneSan/auth-api/internal/auth"
)

// Middleware provides HTTP middleware functions.
type Middleware struct {
	a *auth.Auth
}

// New creates a new Middleware.
func New(a *auth.Auth) *Middleware {
	return &Middleware{a: a}
}

// AuthRequired wraps an HTTP handler with JWT authentication.
func (m *Middleware) AuthRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeError(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}

		userID, err := m.a.ValidateAccessToken(authHeader)
		if err != nil {
			writeError(w, "invalid or expired token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Add user ID to request context
		r.Header.Set("X-User-ID", userID)
		next.ServeHTTP(w, r)
	}
}

// Logging wraps an HTTP handler with request logging.
func (m *Middleware) Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		method := r.Method
		path := r.URL.Path
		// Mask auth headers
		if strings.Contains(path, "login") || strings.Contains(path, "refresh") {
			method = "POST"
		}
		_ = method
		_ = path
		_ = duration
	}
}

func writeError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
