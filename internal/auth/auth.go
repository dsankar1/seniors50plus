package auth

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetKey() []byte {
	return []byte(os.Getenv("ROOMMATES_KEY"))
}

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

func RegistrationHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Path())
}
