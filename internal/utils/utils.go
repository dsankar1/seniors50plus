package utils

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetIdFromContext(c echo.Context, id *uint) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		*id = uint(token.Claims.(jwt.MapClaims)["id"].(float64))
		return nil
	} else {
		return errors.New("Token type assertion failed")
	}
}
