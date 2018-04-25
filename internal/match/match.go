package match

import (
	"net/http"
	"seniors50plus/internal/models"
	"seniors50plus/internal/utils"

	"github.com/labstack/echo"
)

func FindMatchesHandler(c echo.Context) error {
	var token models.Token
	if err := utils.GetTokenFromContext(c, &token); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error connecting to database")
	}
	defer dbc.Close()
	var req models.MatchRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	elastic := models.NewElasticClient()
	results := models.QueryResponse{}
	if err := elastic.Get(&req, &results); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	filtered := []models.Hit{}
	for _, hit := range results.Hits.Hits {
		if hit.Source.UploaderID != token.ID && hit.Source.AcceptedResidentCount != hit.Source.TargetResidentCount && hit.Source.Budget <= req.BudgetMax && hit.Source.Budget >= req.BudgetMin {
			filtered = append(filtered, hit)
		}
	}
	return c.JSON(http.StatusOK, filtered)
}
