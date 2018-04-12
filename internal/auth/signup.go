package auth

import (
	"errors"
	"net/http"
	"seniors50plus/internal/email"
	"seniors50plus/internal/models"
	"time"
	"unicode"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func EmailConfirmationHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"]

		query := "update users set active=true where email=?"
		dbc := models.NewDatabaseConnection()
		if _, err := dbc.ExecuteQuery(query, email); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.Redirect(http.StatusPermanentRedirect, "/home")
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token type assertion failed?")
	}
}

func SignupHandler(c echo.Context) error {
	var req models.SignupRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Empty request body")
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing/Invalid fields")
	}
	if err := validateRequest(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	passwordHash, _ := HashPassword(req.Password)
	registrationDate := time.Now().String()
	adminLevel := models.AdminLevelUser

	user := models.User{
		Email:            req.Email,
		Firstname:        req.Firstname,
		Lastname:         req.Lastname,
		Gender:           req.Gender,
		PasswordHash:     passwordHash,
		Birthdate:        req.Birthdate,
		AdminLevel:       adminLevel,
		RegistrationDate: registrationDate,
	}

	dbc := models.NewDatabaseConnection()
	if err := dbc.CreateUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = req.Email
		claims["exp"] = time.Now().Add(time.Hour).Unix()
		tokenString, _ := token.SignedString(GetKey())
		tokenString = "https://roommates40plus.com/api/auth/confirmation?token=" + tokenString

		tmpInfo := models.TemplateInfo{
			Firstname: req.Firstname,
			URL:       tokenString,
		}

		if err := email.SendEmail(
			"smtp.gmail.com",
			587,
			"roommates40plus@gmail.com",
			"capst0ne!40Plus",
			[]string{req.Email},
			"Registration Confirmation",
			registrationTemplate,
			tmpInfo,
		); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
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
