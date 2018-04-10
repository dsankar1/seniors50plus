package auth

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

const (
	requiredAge          = 40
	missingFields        = "Missing fields"
	invalidFormat        = "Invalid email format"
	unresolvedHost       = "Unresolved email host"
	nonexistentUser      = "Provided email account doesn't exist"
	registrationTemplate = "../../templates/registration.html"
)

func GetKey() []byte {
	return []byte(os.Getenv("ROOMMATES_KEY"))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
