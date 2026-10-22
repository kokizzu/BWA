package auth

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
)

func TestTokenSigningAndValidation(t *testing.T) {
	service := NewService()
	encoded, err := service.GenerateToken(42)
	if err != nil {
		t.Fatal(err)
	}
	token, err := service.ValidateToken(encoded)
	if err != nil || token == nil || !token.Valid {
		t.Fatalf("generated token rejected: %v", err)
	}
	raw, err := json.Marshal(token.Claims)
	if err != nil {
		t.Fatal(err)
	}
	var claims map[string]any
	if err := json.Unmarshal(raw, &claims); err != nil {
		t.Fatal(err)
	}
	if claims["user_id"] != float64(42) {
		t.Fatalf("user identity changed: %v", claims["user_id"])
	}
	parts := strings.Split(encoded, ".")
	parts[1] = base64.RawURLEncoding.EncodeToString([]byte("{\"user_id\":43}"))
	for name, value := range map[string]string{
		"modified identity": strings.Join(parts, "."),
		"unsigned token":    "eyJhbGciOiJub25lIn0.eyJ1c2VyX2lkIjoxfQ.",
	} {
		t.Run(name, func(t *testing.T) {
			if token, err := service.ValidateToken(value); err == nil || (token != nil && token.Valid) {
				t.Fatal("untrusted token accepted")
			}
		})
	}
}
