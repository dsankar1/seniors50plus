package management

import (
	"net/http"
	"seniors50plus/internal/models"
	"seniors50plus/internal/utils"
	"strconv"

	"github.com/labstack/echo"
)

func ReportUserHandler(c echo.Context) error {
	var req models.ReportRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
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
	user := models.User{
		ID: uint(userId),
	}
	if err := dbc.GetUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	report := models.Report{
		UserID:         token.ID,
		ReportedUserID: user.ID,
		Message:        req.Message,
	}
	if err := dbc.CreateReport(&report); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating report")
	}
	return c.JSON(http.StatusOK, report)
}

func ResolveReportsHandler(c echo.Context) error {
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
	if err := dbc.DeleteReports(uint(userId)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	res := struct {
		Message string
	}{
		"Reports Deleted",
	}
	return c.JSON(http.StatusOK, res)
}

func GetReportsHandler(c echo.Context) error {
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
	var reports []models.Report
	if err := dbc.GetAllReports(&reports); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, reports)
}
