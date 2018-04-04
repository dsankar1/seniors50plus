package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
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
	info := new(RegistrationRequest)
	if err := c.Bind(info); err != nil {
		return err
	}
	fmt.Println("Received Info:", info)
	if !Validate(info) {
		return errors.New("Missing fields")
	}
	return c.JSON(http.StatusOK, c.Path())
}

func Validate(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
