package auth

import (
	"net/http"

	"github.com/labstack/echo"
)

func AuthenticationHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Path())
}

func RegistrationHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Path())
}
