package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func GenerateToken() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", sha256.Sum256([]byte(base64.StdEncoding.EncodeToString(randomBytes)[:32])))
}
