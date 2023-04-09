package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"net/mail"
	"strings"
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

func SetSessionNumInCookie(w http.ResponseWriter, num string) {
	cookie := http.Cookie{
		Name:     "sessionNum",
		Value:    num,
		Expires:  time.Now().Add(3600 * 24 * 3 * time.Second), // 3 days
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode, // TODO: change this to Strict (maybe)
	}
	http.SetCookie(w, &cookie)
}

func GetSessionNumFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("sessionNum")
	if err != nil {
		// when there is no cookie set to the browser
		return ""
	}
	return cookie.Value
}

// SetSessionCookie sets a session cookie
func SetSessionCookie(w http.ResponseWriter, no string, sid string) {
	cookie := http.Cookie{
		Name:     "sessionid" + no,
		Value:    sid,
		Expires:  time.Now().Add(3600 * 24 * 3 * time.Second), // 3 days
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode, // TODO: change this to Strict (maybe)
	}
	http.SetCookie(w, &cookie)
}

// GetSessionCookie obtains sessionID inside the cookie
func GetSessionCookie(r *http.Request, no string) string {
	cookie, err := r.Cookie("sessionid" + no)
	if err != nil {
		// when there is no cookie set to the browser
		return ""
	}
	return cookie.Value
}

// SetUsernameCookie sets username cookie
func SetUsernameCookie(w http.ResponseWriter, username string) {
	usernameCookie := http.Cookie{
		Name:     "username",
		Value:    username,
		Expires:  time.Now().Add(3600 * 24 * 3 * time.Second), // 3 days
		HttpOnly: false,                                       // so that the browser's javascript extract data from the username cookie
		Secure:   true,
		SameSite: http.SameSiteLaxMode, // TODO: change this to Strict (maybe)
	}
	http.SetCookie(w, &usernameCookie)
}

func GetUsernameCookie(r *http.Request) string {
	cookie, err := r.Cookie("username")
	if err != nil {
		// when there is no cookie set to the browser
		return ""
	}
	return cookie.Value
}

// SetEmailCookie sets email cookie
func SetEmailCookie(w http.ResponseWriter, email string) {
	cookie := http.Cookie{
		Name:     "email",
		Value:    email,
		Expires:  time.Now().Add(3600 * 24 * 1 * time.Second), // 3 days
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode, // TODO: change this to Strict (maybe)
	}
	http.SetCookie(w, &cookie)
}

// DeleteCookie deletes the named cookie from browser
func DeleteCookie(w http.ResponseWriter, key string, value string) {
	cookie := &http.Cookie{
		Name:     key,
		Value:    value,
		Expires:  time.Now(),
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
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
func ValidateSignupInput(email string, password string) string {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "email is invalid"
	}
	// TODO: fix validations regex

	password = strings.ToLower(password)

	// valid
	return ""
}

func ValidateEmail(email string) string {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "email is invalid"
	}
	return ""
}
