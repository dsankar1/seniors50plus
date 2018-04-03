package middleware

import (
	"seniors50plus/internal/auth"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ApplyMiddleware(e *echo.Echo) {

	// Redirects to https
	//e.Pre(middleware.HTTPSRedirect())

	// Checks incoming requests to api endpoints for JWT (excludes authenticate and register endpoints)
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: auth.GetKey(),
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/api") {
				if c.Path() == "/api/authenticate" || c.Path() == "/api/register" || c.Path() == "/api/test/authenticate" {
					return true
				}
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
