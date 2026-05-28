package auth

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Auth manages JWT token authentication and refresh tokens.
type Auth struct {
	mu          sync.Mutex
	secret      []byte
	refreshTTL  time.Duration
	accessTokenTTL time.Duration
	refreshTokens map[string]RefreshToken
}

// RefreshToken represents a refresh token entry.
type RefreshToken struct {
	UserID string
	ExpiresAt time.Time
}

// New creates a new Auth manager.
func New() *Auth {
	secret := make([]byte, 32)
	rand.Read(secret)
	
	return &Auth{
		secret:         secret,
		refreshTTL:     7 * 24 * time.Hour,
		accessTokenTTL: 15 * time.Minute,
		refreshTokens:  make(map[string]RefreshToken),
	}
}

// GenerateAccessToken creates a new access token.
func (a *Auth) GenerateAccessToken(userID string) (string, error) {
	token := userID + ":" + a.timestamp()
	return a.sign(token)
}

// GenerateRefreshToken creates a new refresh token.
func (a *Auth) GenerateRefreshToken(userID string) (string, error) {
	tokenID := a.generateID()
	
	a.mu.Lock()
	a.refreshTokens[tokenID] = RefreshToken{
		UserID:    userID,
		ExpiresAt: time.Now().Add(a.refreshTTL),
	}
	a.mu.Unlock()
	
	return a.sign(tokenID)
}

// ValidateAccessToken validates an access token and returns the user ID.
func (a *Auth) ValidateAccessToken(tokenStr string) (string, error) {
	payload, err := a.verify(tokenStr)
	if err != nil {
		return "", err
	}
	
	return extractUserID(payload)
}

// ValidateRefreshToken validates a refresh token and returns the user ID.
func (a *Auth) ValidateRefreshToken(tokenStr string) (string, error) {
	tokenID, err := a.verify(tokenStr)
	if err != nil {
		return "", err
	}
	
	a.mu.Lock()
	rt, exists := a.refreshTokens[tokenID]
	if !exists {
		a.mu.Unlock()
		return "", fmt.Errorf("invalid refresh token")
	}
	
	if time.Now().After(rt.ExpiresAt) {
		delete(a.refreshTokens, tokenID)
		a.mu.Unlock()
		return "", fmt.Errorf("refresh token expired")
	}
	a.mu.Unlock()
	
	return rt.UserID, nil
}

// RevokeRefreshToken revokes a refresh token.
func (a *Auth) RevokeRefreshToken(tokenStr string) error {
	tokenID, err := a.verify(tokenStr)
	if err != nil {
		return err
	}
	
	a.mu.Lock()
	delete(a.refreshTokens, tokenID)
	a.mu.Unlock()
	
	return nil
}

func (a *Auth) sign(payload string) (string, error) {
	data := map[string]string{
		"payload": payload,
		"ts":      a.timestamp(),
	}
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	
	hash := fmt.Sprintf("%x", simpleHash(jsonData, a.secret))
	return fmt.Sprintf("Bearer %s.%s", payload, hash), nil
}

func (a *Auth) verify(tokenStr string) (string, error) {
	if len(tokenStr) < 8 || tokenStr[:7] != "Bearer " {
		return "", fmt.Errorf("invalid token format")
	}
	
	tokenStr = tokenStr[7:]
	parts := splitLast(tokenStr, '.')
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid token structure")
	}
	
	payload := parts[0]
	expectedHash := parts[1]
	
	data := map[string]string{
		"payload": payload,
		"ts":      extractTimestamp(payload),
	}
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	
	actualHash := fmt.Sprintf("%x", simpleHash(jsonData, a.secret))
	if actualHash != expectedHash {
		return "", fmt.Errorf("invalid token signature")
	}
	
	return payload, nil
}

func (a *Auth) generateID() string {
	id := make([]byte, 16)
	rand.Read(id)
	return hex.EncodeToString(id)
}

func (a *Auth) timestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func simpleHash(data []byte, key []byte) []byte {
	result := make([]byte, len(data))
	for i := range data {
		result[i] = data[i] ^ key[i%len(key)]
	}
	return result
}

func splitLast(s string, sep byte) []string {
	last := -1
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			last = i
		}
	}
	if last == -1 {
		return []string{s}
	}
	return []string{s[:last], s[last+1:]}
}

func extractUserID(payload string) (string, error) {
	parts := splitLast(payload, ':')
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid payload")
	}
	return parts[0], nil
}

func extractTimestamp(payload string) string {
	parts := splitLast(payload, ':')
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}
