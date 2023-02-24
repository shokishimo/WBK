package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/mail"
	"regexp"
	"time"
)

// GenerateSessionID generates a sessionID
func GenerateSessionID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	sessionID := hex.EncodeToString(b)
	return sessionID
}

// GeneratePasscode generates a passcode of length 6
func GeneratePasscode() string {
	// initialize the random number generator
	rand.Seed(time.Now().UnixNano())
	// generate a random integer between 0 and 999999 (inclusive)
	randomInt := rand.Intn(1000000)
	// format the integer as a string with leading zeros and length 6
	passcode := fmt.Sprintf("%06d", randomInt)
	return passcode
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
