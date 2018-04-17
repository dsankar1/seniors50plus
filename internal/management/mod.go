package management

import (
	"net/http"
	"seniors50plus/internal/models"
	"seniors50plus/internal/utils"
	"strconv"

	"github.com/labstack/echo"
)

func ModUserHandler(c echo.Context) error {
	var token models.Token
	if err := utils.GetTokenFromContext(c, &token); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if token.Admin != models.AdminLevelSusan {
		return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
	}
	user := models.User{
		ID: uint(userId),
	}
	if err := dbc.GetUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	user.AdminLevel = models.AdminLevelModerator
	if err := dbc.UpdateUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error updating user")
	}
	return c.JSON(http.StatusOK, user)
}

func UnmodUserHandler(c echo.Context) error {
	var token models.Token
	if err := utils.GetTokenFromContext(c, &token); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if token.Admin != models.AdminLevelSusan {
		return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
	}
	user := models.User{
		ID: uint(userId),
	}
	if err := dbc.GetUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	user.AdminLevel = models.AdminLevelUser
	if err := dbc.UpdateUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error updating user")
	}
	return c.JSON(http.StatusOK, user)
}

func GetModsHandler(c echo.Context) error {
	var token models.Token
	if err := utils.GetTokenFromContext(c, &token); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if token.Admin != models.AdminLevelSusan {
		return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
	}
	var mods []models.User
	if err := dbc.GetMods(&mods); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mods)
}
