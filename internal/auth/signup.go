package auth

import (
	"net/http"
	"seniors50plus/internal/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func EmailConfirmationHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		userId := uint(token.Claims.(jwt.MapClaims)["id"].(float64))
		dbc := models.NewDatabaseConnection()
		if err := dbc.Open(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
		}
		defer dbc.Close()
		user := models.User{
			ID: userId,
		}
		if err := dbc.GetUser(&user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		user.Active = true
		if err := dbc.UpdateUser(&user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		res := struct {
			Message string `json:"message"`
		}{
			"Success",
		}
		return c.JSON(http.StatusOK, res)
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token type assertion failed?")
	}
}

func SignupHandler(c echo.Context) error {
	var req models.SignupRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := validateRequest(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	passwordHash, _ := HashPassword(req.Password)
	user := models.User{
		Email:        req.Email,
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
		Gender:       req.Gender,
		Birthdate:    req.Birthdate,
		PasswordHash: passwordHash,
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if err := dbc.CreateUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := SendConfirmationEmail(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res := struct {
		Message string `json:"message"`
	}{
		"Confirmation email sent to " + user.Email,
	}
	return c.JSON(http.StatusOK, res)
}
