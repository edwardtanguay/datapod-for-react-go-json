package utils

import (
	"crypto/rand"
	"math/big"
)

// GenerateShortUUID generates a short unique identifier
// length specifies the number of characters to generate
func GenerateShortUUID(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		randomByte, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err) // Handle appropriately in production
		}
		bytes[i] = charset[randomByte.Int64()]
	}
	return string(bytes)
}
