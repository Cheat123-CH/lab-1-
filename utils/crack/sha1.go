package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

// HashSHA1 returns the lowercase hex SHA-1 of the input string.
func HashSHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
