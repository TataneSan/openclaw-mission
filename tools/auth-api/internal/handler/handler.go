package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/TataneSan/auth-api/internal/auth"
)

// Handler handles HTTP requests for authentication endpoints.
type Handler struct {
	a *auth.Auth
}

// New creates a new Handler.
func New(a *auth.Auth) *Handler {
	return &Handler{a: a}
}

// Health returns a simple health check response.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Login handles user login and returns access and refresh tokens.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, "failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req struct {
		UserID string `json:"user_id"`
		Secret string `json:"secret"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		writeError(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.UserID == "" || req.Secret == "" {
		writeError(w, "fields 'user_id' and 'secret' are required", http.StatusBadRequest)
		return
	}

	// Simple authentication (in production, use a proper user store)
	if req.UserID != req.Secret {
		writeError(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, err := h.a.GenerateAccessToken(req.UserID)
	if err != nil {
		writeError(w, "failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := h.a.GenerateRefreshToken(req.UserID)
	if err != nil {
		writeError(w, "failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Refresh handles token refresh.
func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, "failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		writeError(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.RefreshToken == "" {
		writeError(w, "field 'refresh_token' is required", http.StatusBadRequest)
		return
	}

	userID, err := h.a.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		writeError(w, "invalid refresh token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Revoke old refresh token
	h.a.RevokeRefreshToken(req.RefreshToken)

	accessToken, err := h.a.GenerateAccessToken(userID)
	if err != nil {
		writeError(w, "failed to generate access token", http.StatusInternalServerError)
		return
	}

	newRefreshToken, err := h.a.GenerateRefreshToken(userID)
	if err != nil {
		writeError(w, "failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": newRefreshToken,
	})
}

// Logout handles token revocation.
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, "failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		writeError(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.RefreshToken == "" {
		writeError(w, "field 'refresh_token' is required", http.StatusBadRequest)
		return
	}

	if err := h.a.RevokeRefreshToken(req.RefreshToken); err != nil {
		writeError(w, "failed to revoke token: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "token revoked"})
}

func writeError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
