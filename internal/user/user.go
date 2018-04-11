package user

import (
	"net/http"
	"seniors50plus/internal/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func ModifyUserHandler(c echo.Context) error {
	var userUpdate models.User
	if err := c.Bind(&userUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Empty request body")
	}
	if err := c.Validate(&userUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing/Invalid fields")
	}
	if token, ok := c.Get("user").(*jwt.Token); ok {
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		userUpdate.Email = email
		dbc := models.NewDatabaseConnection()
		if active, err := dbc.IsActive(email); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else if !active {
			return echo.NewHTTPError(http.StatusBadRequest, "Account not activated")
		}
		if user, err := dbc.EditUser(&userUpdate); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			return c.JSON(http.StatusOK, user)
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Token type assertion failed?")
	}
}
