package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	hashTime      = 1
	memory        = 64 * 1024
	threads       = 4
	keyLength     = 32
)

func HashPassword(password string) (string, error) {

	// Generate salt
	salt := make([]byte, 16) // 16 byte salt
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Hash password
	hash := argon2.IDKey([]byte(password), salt, uint32(hashTime), uint32(memory), uint8(threads), uint32(keyLength))

	// Encode salt and hash to base64
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	// Return encoded salt and hash concatenated with '$'
	return encodedSalt + "$" + encodedHash, nil
}

func ComparePassword(password, hash string) (bool, error) {
	// Split hash string into encoded salt and hash
	parts := strings.Split(hash, "$")
	if len(parts) != 2 {
		return false, errors.New("invalid hash format")
	}
	encodedSalt := parts[0]
	encodedHash := parts[1]

	// Decode salt and hash from base64
	salt, err := base64.RawStdEncoding.DecodeString(encodedSalt)
	if err != nil {
		return false, err
	}
	decodedHash, err := base64.RawStdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false, err
	}

	// Hash password with retrieved salt
	newHash := argon2.IDKey([]byte(password), salt, uint32(hashTime), uint32(memory), uint8(threads), uint32(keyLength))

	// Compare the newly hashed password with the original hash
	return subtle.ConstantTimeCompare(decodedHash, newHash) == 1, nil
}
