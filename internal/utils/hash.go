package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSha256ToHex(input string) string {
	hashOperator := sha256.New()
	hashOperator.Write([]byte(input))
	hash := hex.EncodeToString(hashOperator.Sum(nil))
	return hash
}
