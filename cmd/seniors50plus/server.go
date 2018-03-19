package main

import (
	"net/http"
	"seniors50plus/internal/auth"
	"seniors50plus/internal/middleware"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	//e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("daryans-page.com")
	//e.AutoTLSManager.Cache = autocert.DirCache("./var/www/.cache")
	middleware.ApplyMiddleware(e)

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ message string }{"Welcome to the API."})
	})
	e.POST("/api/authenticate", auth.AuthenticationHandler)
	e.POST("/api/register", auth.RegistrationHandler)

	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	//e.Logger.Fatal(e.StartAutoTLS(":443"))
	e.Logger.Fatal(e.StartServer(s))
}
