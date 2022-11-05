package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	saltSize = 32
)

func GenerateRandomSalt(saltSize int) ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		return nil, fmt.Errorf("Faied generate random salt. error: %v", err)
	}

	return salt, nil

}

//Combine password and salt then hash them using SHA-512
//algorithm and then return the hashed password as a hex string
func PasswordHash(password string, salt []byte) (string, error) {

	var passwordBytes = []byte(password)

	var sha512Hasher = sha256.New()

	passwordBytes = append(passwordBytes, salt...)

	_, err := sha512Hasher.Write(passwordBytes)

	if err != nil {
		return "", fmt.Errorf("failed to write passwordBytes to sha256")
	}

	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex, nil

}
