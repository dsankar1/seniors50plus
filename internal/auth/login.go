package auth

import (
	"net/http"
	"seniors50plus/internal/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func LoginHandler(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Empty request body")
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing/Invalid fields")
	}
	dbc := models.NewDatabaseConnection()
	user := models.User{
		Email: req.Email,
	}
	if err := dbc.QueryUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := dbc.AttachTags(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if CheckPasswordHash(req.Password, user.PasswordHash) == false {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect password")
	}
	if user.Active == false {
		return echo.NewHTTPError(http.StatusBadRequest, "Account not activated")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = req.Email
	claims["admin"] = user.AdminLevel
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
	tokenString, _ := token.SignedString(GetKey())
	return c.JSON(http.StatusOK, models.LoginResponse{User: &user, Token: tokenString})
}
