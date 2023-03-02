package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
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

// SetSessionCookie sets a session cookie
func SetSessionCookie(w http.ResponseWriter, sid string) {
	cookie := http.Cookie{
		Name:     "sessionid",
		Value:    sid,
		Expires:  time.Now().Add(3600 * 24 * 3 * time.Second), // 3 days
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode, // TODO: change this to Strict (maybe)
	}
	http.SetCookie(w, &cookie)
}

// GetSessionCookie obtains sessionID inside the cookie
func GetSessionCookie(r *http.Request) string {
	cookie, err := r.Cookie("sessionid")
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
func ValidateSignupInput(email string, password string) bool {
	// TODO: fix validations regex
	//_, err := mail.ParseAddress(email)
	//if err != nil {
	//	return false
	//}
	//regex, err := regexp.Compile(`^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[\\^$.|?*+\-\[\]{}()]).{8,32}$`)
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return false
	//}
	//if regex.FindString(password) == "" {
	//	return false
	//}
	return true
}
