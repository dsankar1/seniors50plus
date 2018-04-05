package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/badoux/checkmail"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const (
	requiredAge     = 40
	missingFields   = "Missing fields"
	invalidFormat   = "Invalid email format"
	unresolvedHost  = "Unresolved email host"
	nonexistentUser = "Provided email account doesn't exist"
)

func GetKey() []byte {
	return []byte(os.Getenv("ROOMMATES_KEY"))
}

func AuthenticationHandler(c echo.Context) error {
	creds := new(AuthRequest)
	if err := c.Bind(creds); err != nil {
		return err
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = creds.Email
	claims["created"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()

	tokenString, _ := token.SignedString(GetKey())
	return c.JSON(http.StatusOK, tokenString)
}

func RegistrationHandler(c echo.Context) error {
	info := new(RegistrationRequest)
	if err := c.Bind(info); err != nil {
		return err
	}
	fmt.Println("Received Info:", info)
	if err2 := validateRegistration(info); err2 != nil {
		return c.JSON(http.StatusBadRequest, Error{err2.Error()})
	}

	SendEmail(
		"smtp.gmail.com",
		587,
		"roommates40plus@gmail.com",
		"capst0ne!40Plus",
		[]string{"daryan.sankar1@gmail.com"},
		"Registration Confirmation",
		info,
	)

	return c.JSON(http.StatusOK, c.Path())
}

func validateRegistration(info *RegistrationRequest) error {
	if err := checkRegistrationFields(info); err != nil {
		return err
	}
	if err2 := validateAge(info.Birthdate); err2 != nil {
		return err2
	}
	if err3 := validateEmail(info.Email); err3 != nil {
		return err3
	}
	return nil
}

func validateAge(birthdate string) error {
	format := "2006-01-02"
	start, _ := time.Parse(format, birthdate)
	end := time.Since(start)
	years := int((end / time.Hour / 24 / 365).Nanoseconds())
	//fmt.Printf("Age: %d years", years)
	if years < requiredAge {
		return errors.New("Must be 40 or older to create an account")
	}
	return nil
}

func checkRegistrationFields(info *RegistrationRequest) error {
	fieldsExist := info.Email != "" && info.Password != "" && info.Firstname != "" && info.Lastname != "" && info.Birthdate != ""
	if !fieldsExist {
		return errors.New(missingFields)
	}
	return nil
}

func validateEmail(email string) error {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return errors.New(invalidFormat)
	}
	return nil
}
