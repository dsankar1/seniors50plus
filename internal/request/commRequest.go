package request

import (
	"net/http"
	"seniors50plus/internal/models"
	"seniors50plus/internal/utils"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func CreateCommunicationRequestHandler(c echo.Context) error {
	var tokenId uint
	if err := utils.GetIdFromContext(c, &tokenId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	offerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid offer ID")
	}
	offer := models.RoommateOffer{
		ID: uint(offerId),
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.GetOffer(&offer); err != nil {
		return err
	}
	if offer.UploaderID == tokenId {
		echo.NewHTTPError(http.StatusBadRequest, "Attempted to request your own post")
	}
	if offer.AcceptedResidentCount == offer.TargetResidentCount {
		echo.NewHTTPError(http.StatusConflict, "Offer is already full")
	}
	request := models.Request{
		UserID:  tokenId,
		OfferID: uint(offerId),
	}
	if err := dbc.CreateCommunicationRequest(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, request)
}

func DeleteCommunicationRequestHandler(c echo.Context) error {
	var tokenId uint
	if err := utils.GetIdFromContext(c, &tokenId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	offerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid offer ID")
	}
	request := models.Request{
		UserID:  tokenId,
		OfferID: uint(offerId),
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.DeleteCommunicationRequest(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, request)
}

func RespondToCommunicationRequestHandler(c echo.Context) error {
	var tokenId uint
	if err := utils.GetIdFromContext(c, &tokenId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	requestId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid offer ID")
	}
	status := strings.ToLower(c.QueryParam("status"))
	if status != models.RequestStatusAccepted && status != models.RequestStatusDenied {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid status")
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
	if offer.UploaderID != tokenId {
		return echo.NewHTTPError(http.StatusUnauthorized, "Offer uploader ID doesn't match token ID")
	}
	request.Status = status
	if err := dbc.UpdateCommunicationRequest(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, request)
}
