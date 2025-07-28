package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"

	"golang.org/x/crypto/pbkdf2"
)

const (
	djangoPBKDF2Iterations = 1000000
	djangoPBKDF2KeyLength = 64
	djangoPBKDF2SaltLength = 12
	alphanumericChars      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateDjangoCompatibleSalt() (string, error) {
	salt := make([]byte, djangoPBKDF2SaltLength)
	for i := 0; i < djangoPBKDF2SaltLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphanumericChars))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number for salt: %w", err)
		}
		salt[i] = alphanumericChars[num.Int64()]
	}
	return string(salt), nil
}

func HashPasswordForDjango(rawPassword string) (string, error) {
	salt, err := GenerateDjangoCompatibleSalt()
	if err != nil {
		return "", fmt.Errorf("could not generate salt: %w", err)
	}

	hashBytes := pbkdf2.Key(
		[]byte(rawPassword),
		[]byte(salt), 
		djangoPBKDF2Iterations,
		djangoPBKDF2KeyLength,
		sha256.New, 
	)

	encodedHash := base64.StdEncoding.EncodeToString(hashBytes)

	return fmt.Sprintf("pbkdf2_sha256$%d$%s$%s", djangoPBKDF2Iterations, salt, encodedHash), nil
}

