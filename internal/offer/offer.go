package offer

import (
	"net/http"
	"seniors50plus/internal/models"
	"seniors50plus/internal/utils"
	"strconv"

	"github.com/jinzhu/gorm"

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
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if err := dbc.GetOffer(&offer); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusOK, []models.Request{})
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := dbc.AttachResidents(&offer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	filtered := []models.Request{}
	for i := 0; i < len(offer.Residents); i++ {
		if offer.Residents[i].Status == models.RequestStatusAccepted {
			filtered = append(filtered, offer.Residents[i])
		}
	}
	offer.Residents = filtered
	return c.JSON(http.StatusOK, offer)
}

func GetOfferEmailHandler(c echo.Context) error {
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
	var res struct {
		Email string
	}

	request := models.Request{
		UserID:  token.ID,
		OfferID: uint(offerId),
	}
	if err := dbc.GetCommunicationRequest(&request); err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if request.Status != models.RequestStatusAccepted {
		return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized to view this email")
	}
	offer := models.RoommateOffer{
		ID: uint(offerId),
	}
	if err := dbc.GetOffer(&offer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	user := models.User{
		ID: offer.UploaderID,
	}
	if err := dbc.GetUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	res.Email = user.Email
	return c.JSON(http.StatusOK, res)
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
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if err := dbc.GetOffer(&offer); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusOK, struct{}{})
		}
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
		if err := dbc.Open(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
		}
		defer dbc.Close()
		offer.UploaderID = userId
		if err := dbc.CreateOffer(&offer); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		offer.Requests = []models.Request{}
		offer.Residents = []models.Request{}
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
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	if err := dbc.DeleteOffer(&offer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	dbc.Close()
	return c.JSON(http.StatusOK, offer)
}
