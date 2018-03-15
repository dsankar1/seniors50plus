package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
}

func main() {
	e := echo.New()

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/api/") || c.Path() == "/api/authenticate" || c.Path() == "/api/register" {
				return true
			}
			return false
		},
	}))

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: true,
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/api/") {
				return true
			}
			return false
		},
	}))

	e.GET("/api/authenticate", func(c echo.Context) error {
		//user := User{c.Param("name")}
		return c.JSON(http.StatusOK, c.Path())
	})

	e.Start(":1323")
}
