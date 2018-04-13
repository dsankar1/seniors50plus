package auth

import (
	"net/http"
	"seniors50plus/internal/email"
	"seniors50plus/internal/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func EmailConfirmationHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		email := token.Claims.(jwt.MapClaims)["email"].(string)

		dbc := models.NewDatabaseConnection()
		user := models.User{
			Email:  email,
			Active: true,
		}
		if err := dbc.UpdateUser(&user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.Redirect(http.StatusPermanentRedirect, "/home")
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token type assertion failed?")
	}
}

func SignupHandler(c echo.Context) error {
	var req models.SignupRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Empty request body")
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing/Invalid fields")
	}
	if err := validateRequest(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	passwordHash, _ := HashPassword(req.Password)
	user := models.User{
		Email:        req.Email,
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
		Gender:       req.Gender,
		Birthdate:    req.Birthdate,
		PasswordHash: passwordHash,
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.CreateUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenString, _ := token.SignedString(GetKey())
	tokenString = "https://roommates40plus.com/api/auth/confirmation?token=" + tokenString

	tmpInfo := models.TemplateInfo{
		Firstname: req.Firstname,
		URL:       tokenString,
	}

	if err := email.SendEmail(
		"smtp.gmail.com",
		587,
		"roommates40plus@gmail.com",
		"capst0ne!40Plus",
		[]string{req.Email},
		"Registration Confirmation",
		registrationTemplate,
		tmpInfo,
	); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
