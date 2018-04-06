package auth

import (
	"errors"
	"fmt"
	"net/http"
	"seniors50plus/internal/database"
	"seniors50plus/internal/user"
	"time"
	"unicode"

	"github.com/badoux/checkmail"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func RegistrationConfirmationHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		fmt.Println(token)
		claims := token.Claims.(jwt.MapClaims)

		email := claims["email"]
		password := claims["password"]
		firstname := claims["firstname"]
		lastname := claims["lastname"]
		birthdate := claims["birthdate"]
		gender := claims["gender"]
		registrationDate := time.Now().String()
		adminLevel := user.AdminLevelUser

		query := fmt.Sprintf(`insert into users (email, first_name, last_name, gender, birthdate, admin_level, password, registration_date)
			values ('%v','%v','%v','%v','%v','%v','%v','%v')`, email, firstname, lastname, gender, birthdate, adminLevel, password, registrationDate)
		fmt.Println("Query:", query)
		results, err := database.ExecuteQuery(query)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Error{err.Error()})
		}
		fmt.Println(results)
		return c.JSON(http.StatusOK, token)
	} else {
		return c.JSON(http.StatusBadRequest, Error{Message: "Type assertion for JWT failed"})
	}
}

func RegistrationHandler(c echo.Context) error {
	info := new(RegistrationRequest)
	if err := c.Bind(info); err != nil {
		return err
	}
	fmt.Println("Received Info:", info)
	if err := validateRegistration(info); err != nil {
		return c.JSON(http.StatusBadRequest, Error{err.Error()})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	passwordHash, _ := HashPassword(info.Password)

	claims["email"] = info.Email
	claims["password"] = passwordHash
	claims["firstname"] = info.Firstname
	claims["lastname"] = info.Lastname
	claims["birthdate"] = info.Birthdate
	claims["gender"] = info.Gender
	claims["created"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, _ := token.SignedString(GetKey())
	tokenString = "https://roommates40plus.com/api/register/confirmation?token=" + tokenString

	tmpInfo := TemplateInfo{
		info.Firstname,
		tokenString,
	}

	err := SendEmail(
		"smtp.gmail.com",
		587,
		"roommates40plus@gmail.com",
		"capst0ne!40Plus",
		[]string{info.Email},
		"Registration Confirmation",
		registrationTemplate,
		tmpInfo,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{err.Error()})
	}
	return c.JSON(http.StatusOK, info)
}

func validateRegistration(info *RegistrationRequest) error {
	if err := checkRegistrationFields(info); err != nil {
		return err
	}
	if err := validateName(info.Firstname, info.Lastname); err != nil {
		return err
	}
	if err := validateGender(info.Gender); err != nil {
		return err
	}
	if err := validateAge(info.Birthdate); err != nil {
		return err
	}
	if err := validateEmail(info.Email); err != nil {
		return err
	}
	return nil
}

func validateName(firstname string, lastname string) error {
	if AssertAlphabetic(firstname) && AssertAlphabetic(lastname) {
		return nil
	}
	return errors.New("First or last name contains non-letters")
}

func validateGender(gender string) error {
	if gender == user.Male || gender == user.Female {
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
	//fmt.Printf("Age: %d years", years)
	if years < requiredAge {
		return errors.New("Must be 40 or older to create an account")
	}
	return nil
}

func checkRegistrationFields(info *RegistrationRequest) error {
	fieldsExist := info.Email != "" && info.Password != "" && info.Firstname != "" && info.Lastname != "" && info.Birthdate != "" && info.Gender != ""
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
