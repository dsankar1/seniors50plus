package management

import (
	"net/http"
	"seniors50plus/internal/models"
	"seniors50plus/internal/utils"
	"strconv"

	"github.com/labstack/echo"
)

func GetBannedUsersHandler(c echo.Context) error {
	var token models.Token
	if err := utils.GetTokenFromContext(c, &token); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if token.Admin != models.AdminLevelModerator && token.Admin != models.AdminLevelSusan {
		return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
	}
	var bans []models.Ban
	if err := dbc.GetAllBans(&bans); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, bans)
}

func BanUserHandler(c echo.Context) error {
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
	if token.Admin != models.AdminLevelModerator && token.Admin != models.AdminLevelSusan {
		return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
	}
	ban := models.Ban{
		ModID:    token.ID,
		BannedID: uint(userId),
	}
	if err := dbc.CreateBan(&ban); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ban)
}

func UnbanUserHandler(c echo.Context) error {
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
	if token.Admin != models.AdminLevelModerator && token.Admin != models.AdminLevelSusan {
		return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
	}
	ban := models.Ban{
		ID: uint(userId),
	}
	if err := dbc.DeleteBan(&ban); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ban)
}

func BanOfferHandler(c echo.Context) error {
	var token models.Token
	if err := utils.GetTokenFromContext(c, &token); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	offerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if token.Admin != models.AdminLevelModerator && token.Admin != models.AdminLevelSusan {
		return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
	}
	offer := models.RoommateOffer{
		ID: uint(offerId),
	}
	if err := dbc.DeleteOffer(&offer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, offer)
}
