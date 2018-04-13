package user

import (
	"errors"
	"net/http"
	"seniors50plus/internal/models"
	"time"
	"unicode"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const requiredAge = 40

func ModifyUserHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		email := token.Claims.(jwt.MapClaims)["email"].(string)
		dbc := models.NewDatabaseConnection()
		user := models.User{Email: email}
		if err := dbc.QueryUser(&user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if user.Active == false {
			return echo.NewHTTPError(http.StatusBadRequest, "Account not activated")
		}
		var req models.User
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Empty request body")
		}
		if err := c.Validate(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing/Invalid fields")
		}
		user.Firstname = req.Firstname
		user.Lastname = req.Lastname
		user.Gender = req.Gender
		user.Birthdate = req.Birthdate
		user.About = req.About
		user.Tags = req.Tags

		if err := validateUpdate(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := dbc.UpdateUser(&user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Token type assertion failed?")
	}
}

func validateUpdate(user *models.User) error {
	if err := validateName(user.Firstname, user.Lastname); err != nil {
		return err
	}
	if err := validateGender(user.Gender); err != nil {
		return err
	}
	if err := validateAge(user.Birthdate); err != nil {
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
		return errors.New("Birthdate provided puts you under 40")
	}
	return nil
}
