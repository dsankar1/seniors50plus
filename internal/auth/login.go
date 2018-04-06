package auth

import (
	"net/http"
	"seniors50plus/internal/user"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func AuthenticationHandler(c echo.Context) error {
	creds := new(AuthRequest)
	if err := c.Bind(creds); err != nil {
		return err
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = creds.Email
	claims["created"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()

	tokenString, _ := token.SignedString(GetKey())
	return c.JSON(http.StatusOK, tokenString)
}

// Test endpoint for logging in
func AuthenticationHandlerTest(c echo.Context) error {
	creds := new(AuthRequest)
	if err := c.Bind(creds); err != nil {
		return err
	}

	if creds.Email == "" || creds.Password == "" {
		return c.JSON(http.StatusBadRequest, Error{"Missing fields"})
	}
	user := user.ExampleUser
	user.Email = creds.Email
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = creds.Email
	claims["created"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()

	tokenString, _ := token.SignedString(GetKey())

	res := AuthResponse{
		User:  user,
		Token: tokenString,
	}
	return c.JSON(http.StatusOK, res)
}
