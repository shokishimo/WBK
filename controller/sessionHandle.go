package controller

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/mail"
	"regexp"
)

// GenerateSessionID Generates a sessionID
func GenerateSessionID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	sessionID := hex.EncodeToString(b)
	return sessionID
}

// Hash hashes the input and return the hashed value
func Hash(val string) string {
	// Hash the password
	hash := sha256.Sum256([]byte(val))
	// Encode the hash as a hexadecimal string
	hashed := hex.EncodeToString(hash[:])
	return hashed
}

// ValidateSignupInput validates user input to signup or login
func ValidateSignupInput(email string, password string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	regex, err := regexp.Compile(`^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[\\^$.|?*+\-\[\]{}()]).{8,32}$`)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	if regex.FindString(password) == "" {
		return false
	}
	return true
}
