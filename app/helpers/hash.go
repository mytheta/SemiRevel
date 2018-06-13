package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

func ToHash(password string) string {
	converted := sha256.Sum256([]byte(password))
	return hex.EncodeToString(converted[:])
}
