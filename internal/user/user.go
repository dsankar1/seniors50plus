package user

import (
	"errors"
	"fmt"
	"net/http"
	"seniors50plus/internal/models"
	"strconv"
	"time"
	"unicode"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const requiredAge = 40

func GetUserHandler(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := models.User{
		ID: uint(userId),
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.GetUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := dbc.AttachTags(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user.Email = ""
	return c.JSON(http.StatusOK, user)
}

func GetUserListHandler(c echo.Context) error {
	var list []struct {
		ID uint `json:"id" validate:"required"`
	}
	if err := c.Bind(&list); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var users []models.User
	for i, listItem := range list {
		if listItem.ID == 0 {
			continue
		}
		user := models.User{
			ID: listItem.ID,
		}
		dbc := models.NewDatabaseConnection()
		if err := dbc.GetUser(&user); err != nil {
			errAppend := fmt.Sprintf(" with id=%v (index %v)", listItem.ID, i)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()+errAppend)
		}
		user.Email = ""
		users = append(users, user)
	}
	return c.JSON(http.StatusOK, users)
}

func UpdateUserHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		userId := uint(token.Claims.(jwt.MapClaims)["id"].(float64))
		user := models.User{ID: userId}
		dbc := models.NewDatabaseConnection()
		if err := dbc.GetUser(&user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if user.Active == false {
			return echo.NewHTTPError(http.StatusBadRequest, "Account not activated")
		}
		var update models.User
		if err := c.Bind(&update); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(&update); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		user.Firstname = update.Firstname
		user.Lastname = update.Lastname
		user.Gender = update.Gender
		user.Birthdate = update.Birthdate
		user.About = update.About
		user.Tags = update.Tags

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
