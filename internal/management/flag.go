package management

import (
	"net/http"
	"seniors50plus/internal/models"
	"seniors50plus/internal/utils"
	"strconv"

	"github.com/labstack/echo"
)

func GetFlaggedOffers(c echo.Context) error {
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
	var flags []models.Flag
	if err := dbc.GetAllFlags(&flags); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, flags)
}

func FlagOfferHandler(c echo.Context) error {
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
	offer := models.RoommateOffer{
		ID: uint(offerId),
	}
	if err := dbc.GetOffer(&offer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error finding offer")
	}
	if offer.UploaderID == token.ID {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token ID matches offer uploader ID")
	}
	flag := models.Flag{
		UserID:          token.ID,
		ReportedOfferID: uint(offerId),
	}
	if err := dbc.CreateFlag(&flag); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, flag)
}

func UnflagOfferHandler(c echo.Context) error {
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
	flag := models.Flag{
		UserID:          token.ID,
		ReportedOfferID: uint(offerId),
	}
	if err := dbc.DeleteFlag(&flag); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, flag)
}
