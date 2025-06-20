package config

import (
	"os"
	"strconv"
	"time"
)

// JWT-related configuration
type JWTConfig struct {
	Algorithm string
	Secret    []byte
	ExpiresIn time.Duration
}

// LoadJWTConfig loads JWT configuration from .env
func LoadJWTConfig() *JWTConfig {
	// Default values
	cfg := &JWTConfig{
		Algorithm: "HS256",
		Secret:    []byte("super_secret_xxx"),
		ExpiresIn: 30 * time.Minute,
	}

	// Override from environment variables
	if algo := os.Getenv("JWT_ALGO"); algo != "" {
		cfg.Algorithm = algo
	}

	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		cfg.Secret = []byte(secret)
	}

	if expires := os.Getenv("JWT_EXPIRES_MIN"); expires != "" {
		if min, err := strconv.Atoi(expires); err == nil && min > 0 {
			cfg.ExpiresIn = time.Duration(min) * time.Minute
		}
	}

	return cfg
}
