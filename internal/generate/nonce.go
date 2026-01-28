package generate

import (
	"crypto/rand"
	"encoding/base64"
)

// Nonce creates a new, cryptographically secure, base64-encoded nonce.
// Nonces should be at least 128 bits of entropy (16 bytes) to be secure.
func Nonce() (string, error) {
	// 16 bytes is 128 bits of entropy, which is recommended for security.
	nonceBytes := make([]byte, 16)
	if _, err := rand.Read(nonceBytes); err != nil {
		return "", err
	}
	// Encode to Base64 RawURL format for a URL-safe string without padding.
	nonce := base64.RawURLEncoding.EncodeToString(nonceBytes)
	return nonce, nil
}
