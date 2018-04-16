package offer

import (
	"net/http"
	"seniors50plus/internal/models"
	"seniors50plus/internal/utils"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetOfferHandler(c echo.Context) error {
	offerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid offer ID")
	}
	offer := models.RoommateOffer{
		ID: uint(offerId),
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.GetOffer(&offer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := dbc.AttachResidents(&offer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, offer)
}

func GetMyOfferHandler(c echo.Context) error {
	var tokenId uint
	if err := utils.GetIdFromContext(c, &tokenId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	offer := models.RoommateOffer{
		UploaderID: tokenId,
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.GetOffer(&offer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := dbc.AttachRequests(&offer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := dbc.AttachResidents(&offer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, offer)
}

func PostOfferHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		userId := uint(token.Claims.(jwt.MapClaims)["id"].(float64))
		var offer models.RoommateOffer
		if err := c.Bind(&offer); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(&offer); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		dbc := models.NewDatabaseConnection()
		offer.UploaderID = userId
		if err := dbc.CreateOffer(&offer); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, offer)
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token type assertion failed?")
	}
}

func DeleteMyOfferHandler(c echo.Context) error {
	var tokenId uint
	if err := utils.GetIdFromContext(c, &tokenId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	offer := models.RoommateOffer{
		UploaderID: tokenId,
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.DeleteOffer(&offer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, offer)
}
