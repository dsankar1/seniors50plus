package request

import (
	"net/http"
	"seniors50plus/internal/models"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateResidentRequestHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		myId := uint(token.Claims.(jwt.MapClaims)["id"].(float64))
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if myId == uint(userId) {
			return echo.NewHTTPError(http.StatusBadRequest, "Attempted to request yourself")
		}
		offer := models.RoommateOffer{
			UploaderID: myId,
		}
		dbc := models.NewDatabaseConnection()
		if err := dbc.GetOffer(&offer); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if offer.AcceptedResidentCount == offer.TargetResidentCount {
			return echo.NewHTTPError(http.StatusBadRequest, "Offer is already full")
		}
		request := models.Request{
			UserID:  uint(userId),
			OfferID: offer.ID,
		}
		if err := dbc.CreateResidentRequest(&request); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, request)
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token type assertion failed?")
	}
}

func DeleteResidentRequestHandler(c echo.Context) error {
	if token, ok := c.Get("user").(*jwt.Token); ok {
		myId := uint(token.Claims.(jwt.MapClaims)["id"].(float64))
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if myId == uint(userId) {
			return echo.NewHTTPError(http.StatusBadRequest, "Attempted to remove yourself")
		}
		offer := models.RoommateOffer{
			UploaderID: myId,
		}
		dbc := models.NewDatabaseConnection()
		if err := dbc.GetOffer(&offer); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		request := models.Request{
			UserID:  uint(userId),
			OfferID: offer.ID,
		}
		if err := dbc.GetResidentRequest(&request); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if request.Status == models.RequestStatusAccepted {
			if err := dbc.DecrementOffer(&offer); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}
		if err := dbc.DeleteResidentRequest(&request); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		res := struct {
			Message string
		}{
			"Deleted",
		}
		return c.JSON(http.StatusOK, res)
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token type assertion failed?")
	}
}

func RespondToResidentRequestHandler(c echo.Context) error {
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
		if err := dbc.GetResidentRequest(&request); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if request.UserID != userId {
			return echo.NewHTTPError(http.StatusUnauthorized, "You arent the authorized user")
		}
		request.Status = c.QueryParam("status")
		if request.Status == models.RequestStatusAccepted {
			offer := models.RoommateOffer{
				ID: request.OfferID,
			}
			if err := dbc.IncrementOffer(&offer); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			/*if offer.AcceptedResidentCount == offer.TargetResidentCount {
				if err := dbc.RemovePendingResidentRequests(&offer); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
			}*/
		}
		if err := dbc.UpdateResidentRequest(&request); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, request)
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token type assertion failed?")
	}
}
