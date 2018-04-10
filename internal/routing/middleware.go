package routing

import (
	"seniors50plus/internal/auth"
	"seniors50plus/internal/models"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	validator "gopkg.in/go-playground/validator.v9"
)

func RegisterMiddleware(e *echo.Echo) {
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Custom request validator
	e.Validator = &models.Validator{Validator: validator.New()}

	// Checks for token in authorization header if an api request is recieved(excluding auth endpoints)
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: auth.GetKey(),
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/api") {
				if strings.HasPrefix(c.Path(), "/api/auth") {
					return true
				}
				return false
			}
			return true
		},
	}))

	// Checks for token in url param for email confirmation and password reset endpoints
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  auth.GetKey(),
		TokenLookup: "query:token",
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/api/auth/confirmation") ||
				strings.HasPrefix(c.Path(), "/api/auth/passwordreset") {
				return false
			}
			return true
		},
	}))

	// Responds with contents of static folder for all non api requests
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "../../public",
		HTML5:  true,
		Browse: true,
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/api") {
				return true
			}
			return false
		},
	}))

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"https://roommates40plus.com"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
}
