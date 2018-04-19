package request

import (
	"net/http"
	"seniors50plus/internal/models"
	"seniors50plus/internal/utils"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func CreateResidentRequestHandler(c echo.Context) error {
	var tokenId uint
	if err := utils.GetIdFromContext(c, &tokenId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	if uint(userId) == tokenId {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID matches token ID")
	}
	offer := models.RoommateOffer{
		UploaderID: tokenId,
	}
	user := models.User{
		ID: uint(userId),
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if err := dbc.GetUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error finding requested user")
	}
	if err := dbc.GetOffer(&offer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error finding your offer")
	}
	if offer.AcceptedResidentCount == offer.TargetResidentCount {
		return echo.NewHTTPError(http.StatusBadRequest, "Offer is already full")
	}
	request := models.Request{
		OfferID: offer.ID,
		UserID:  uint(user.ID),
	}
	if err := dbc.CreateResidentRequest(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, request)
}

func DeleteResidentRequestHandler(c echo.Context) error {
	var tokenId uint
	if err := utils.GetIdFromContext(c, &tokenId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	if uint(userId) == tokenId {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID matches token ID")
	}
	offer := models.RoommateOffer{
		UploaderID: tokenId,
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if err := dbc.GetOffer(&offer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error finding your offer")
	}
	request := models.Request{
		OfferID: offer.ID,
		UserID:  uint(userId),
	}
	if err := dbc.GetResidentRequest(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error finding request")
	}
	if request.Status == models.RequestStatusAccepted {
		offer.AcceptedResidentCount--
		if err := dbc.UpdateOffer(&offer); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error updating offer")
		}
	}
	if err := dbc.DeleteResidentRequest(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error deleting request")
	}
	return c.JSON(http.StatusOK, request)
}

func RespondToResidentRequestHandler(c echo.Context) error {
	var tokenId uint
	if err := utils.GetIdFromContext(c, &tokenId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	requestId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request ID")
	}
	status := strings.ToLower(c.QueryParam("status"))
	if status != models.RequestStatusAccepted && status != models.RequestStatusDenied {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid status")
	}
	request := models.Request{
		ID: uint(requestId),
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if err := dbc.GetResidentRequest(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if request.UserID != tokenId {
		return echo.NewHTTPError(http.StatusUnauthorized, "Request user ID doesn't match token ID")
	}
	request.Status = status
	if err := dbc.UpdateResidentRequest(&request); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Error updating request status")
	}
	if status == models.RequestStatusAccepted {
		offer := models.RoommateOffer{
			ID: request.OfferID,
		}
		if err := dbc.GetOffer(&offer); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		offer.AcceptedResidentCount++
		if err := dbc.UpdateOffer(&offer); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error updating offer")
		}
		if offer.AcceptedResidentCount == offer.TargetResidentCount {
			if err := dbc.RemovePendingResidentRequests(&offer); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Error removing pending requests")
			}
		}
	}
	return c.JSON(http.StatusOK, request)
}
