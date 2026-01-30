package generate

import (
	"crypto/hmac"
	"crypto/md5"
)

// HMAC creates an HMAC signature for a given message and secret key.
func HMAC(key, message []byte) []byte {
	// Create a new HMAC hash using the desired hash function and the secret key.
	h := hmac.New(md5.New, key)

	// Write the message to the hash.
	h.Write(message)

	// Get the final HMAC signature as a byte slice.
	signature := h.Sum(nil)

	return signature
	// Encode the signature to a hex string for common use cases (e.g., API headers).
	//return hex.EncodeToString(signature)
}
