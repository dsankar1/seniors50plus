package auth

import (
	"errors"
	"os"
	"seniors50plus/internal/models"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const (
	requiredAge          = 40
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

func validateRequest(req *models.SignupRequest) error {
	if err := validateName(req.Firstname, req.Lastname); err != nil {
		return err
	}
	if err := validateGender(req.Gender); err != nil {
		return err
	}
	if err := validateAge(req.Birthdate); err != nil {
		return err
	}
	return nil
}

func validateName(firstname string, lastname string) error {
	if AssertAlphabetic(firstname) && AssertAlphabetic(lastname) {
		return nil
	}
	return errors.New("Name contains non-letters")
}

func validateGender(gender string) error {
	if gender == models.GenderMale || gender == models.GenderFemale {
		return nil
	}
	return errors.New("Invalid gender")
}

func AssertAlphabetic(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func validateAge(birthdate string) error {
	format := "2006-01-02"
	start, err := time.Parse(format, birthdate)
	if err != nil {
		return errors.New("Incorrect date format")
	}
	end := time.Since(start)
	years := int((end / time.Hour / 24 / 365).Nanoseconds())
	if years < requiredAge {
		return errors.New("Must be 40 or older to create an account")
	}
	return nil
}
