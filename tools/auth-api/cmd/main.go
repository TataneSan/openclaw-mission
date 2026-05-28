package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TataneSan/auth-api/internal/auth"
	"github.com/TataneSan/auth-api/internal/handler"
	"github.com/TataneSan/auth-api/internal/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	a := auth.New()
	h := handler.New(a)
	mw := middleware.New(a)

	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("GET /health", handler.Health)

	// Auth endpoints
	mux.HandleFunc("POST /api/auth/login", mw.Logging(h.Login))
	mux.HandleFunc("POST /api/auth/refresh", mw.Logging(h.Refresh))
	mux.HandleFunc("POST /api/auth/logout", mw.Logging(h.Logout))

	log.Printf("auth-api starting on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
