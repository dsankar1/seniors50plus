package user

import (
	"fmt"
	"net/http"
	"seniors50plus/internal/auth"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func ModifyUserHandler(c echo.Context) error {
	var userUpdate User
	if err := c.Bind(&userUpdate); err != nil {
		return err
	}
	if token, ok := c.Get("user").(*jwt.Token); ok {
		
		return c.JSON(http.StatusOK, struct{})
	} else {
		return c.JSON(http.StatusBadRequest, auth.Error{Message: "Type assertion for JWT failed"})
	}
}
