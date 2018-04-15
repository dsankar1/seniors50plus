package request

import (
	"net/http"
	"seniors50plus/internal/models"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateCommunicationRequestHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		userId := uint(token.Claims.(jwt.MapClaims)["id"].(float64))
		offerId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		request := models.Request{
			UserID:  userId,
			OfferID: uint(offerId),
		}
		dbc := models.NewDatabaseConnection()
		if err := dbc.CreateCommunicationRequest(&request); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, request)
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token type assertion failed?")
	}
}

func DeleteCommunicationRequestHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		userId := uint(token.Claims.(jwt.MapClaims)["id"].(float64))
		offerId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		request := models.Request{
			UserID:  userId,
			OfferID: uint(offerId),
		}
		dbc := models.NewDatabaseConnection()
		if err := dbc.DeleteCommunicationRequest(&request); err != nil {
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

func RespondToCommunicationRequestHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		userId := uint(token.Claims.(jwt.MapClaims)["id"].(float64))
		requestId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		request := models.Request{
			ID: uint(requestId),
		}
		dbc := models.NewDatabaseConnection()
		if err := dbc.GetCommunicationRequest(&request); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		offer := models.RoommateOffer{
			ID: request.OfferID,
		}
		if err := dbc.GetOffer(&offer); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if userId != offer.UploaderID {
			return echo.NewHTTPError(http.StatusUnauthorized, "You arent the owner of the offer")
		}
		if status := c.QueryParam("status"); status != models.RequestStatusAccepted && status != models.RequestStatusDenied {
			return echo.NewHTTPError(http.StatusInternalServerError, "Invalid/Missing status parameter")
		} else {
			request.Status = status
		}
		if err := dbc.UpdateCommunicationRequest(&request); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, request)
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token type assertion failed?")
	}
}
