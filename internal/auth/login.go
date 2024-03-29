package auth

import (
	"net/http"
	"seniors50plus/internal/models"
	"time"

	"github.com/jinzhu/gorm"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func LoginHandler(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := models.User{
		Email: req.Email,
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if err := dbc.GetUser(&user); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusInternalServerError, "Email not recognized")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if CheckPasswordHash(req.Password, user.PasswordHash) == false {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect password")
	}
	if user.Active == false {
		if err := SendConfirmationEmail(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		res := struct {
			Message string `json:"message"`
		}{
			"Account not active. Confirmation email resent to " + user.Email,
		}
		return c.JSON(http.StatusOK, res)
	}
	ban := models.Ban{
		BannedID: user.ID,
	}
	if exists, err := dbc.BanExists(&ban); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else {
		if exists {
			return echo.NewHTTPError(http.StatusBadRequest, "Account banned")
		}
	}
	if err := dbc.AttachTags(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := dbc.AttachCommunicationRequests(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := dbc.AttachResidentInvitations(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := dbc.AttachFlags(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["admin"] = user.AdminLevel
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
	tokenString, _ := token.SignedString(GetKey())
	return c.JSON(http.StatusOK, models.LoginResponse{User: &user, Token: tokenString})
}
