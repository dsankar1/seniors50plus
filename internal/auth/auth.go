package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func getSignature() []byte {
	raw, err := ioutil.ReadFile("./secret.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	signature := string(raw)
	return []byte(signature)
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

	//tokenString, _ := token.SignedString(getSignature())
	return c.JSON(http.StatusOK, token)
}

func RegistrationHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Path())
}
