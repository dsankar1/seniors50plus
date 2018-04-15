package offer

import (
	"net/http"
	"seniors50plus/internal/models"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetOfferHandler(c echo.Context) error {
	var offer models.RoommateOffer
	if idParam := c.Param("id"); idParam != "" {
		offerId, err := strconv.Atoi(idParam)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		offer.ID = uint(offerId)
	} else {
		if token, ok := c.Get("user").(*jwt.Token); ok {
			userId := uint(token.Claims.(jwt.MapClaims)["id"].(float64))
			offer.UploaderID = userId
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, "Token type assertion failed?")
		}
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

func DeleteOfferHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		userId := uint(token.Claims.(jwt.MapClaims)["id"].(float64))
		offer := models.RoommateOffer{
			UploaderID: userId,
		}
		dbc := models.NewDatabaseConnection()
		if err := dbc.DeleteOffer(&offer); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		res := struct {
			message string
		}{
			"Deleted",
		}
		return c.JSON(http.StatusOK, res)
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token type assertion failed?")
	}
}
