package utils

import (
	"errors"
	"seniors50plus/internal/models"

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

func GetTokenFromContext(c echo.Context, t *models.Token) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		t.ID = uint(token.Claims.(jwt.MapClaims)["id"].(float64))
		t.Email = token.Claims.(jwt.MapClaims)["email"].(string)
		t.Admin = token.Claims.(jwt.MapClaims)["admin"].(string)
		return nil
	} else {
		return errors.New("Token type assertion failed")
	}
}
