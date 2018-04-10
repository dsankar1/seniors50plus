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
	if user, err := dbc.GetUser(req.Email); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		if user == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Email not recognized")
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
		return c.JSON(http.StatusOK, models.LoginResponse{User: user, Token: tokenString})
	}
}
