package util

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomString generates a random string of the specified length
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}
