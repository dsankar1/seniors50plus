package main

import (
	"net/http"
	"seniors50plus/internal/auth"
	"seniors50plus/internal/match"
	"seniors50plus/internal/middleware"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("roommates40plus.com")
	e.AutoTLSManager.Cache = autocert.DirCache("./var/www/.cache")
	middleware.ApplyMiddleware(e)

	e.POST("/api/authenticate", auth.AuthenticationHandler)

	e.POST("/api/test/authenticate", auth.AuthenticationHandlerTest)

	e.POST("/api/register", auth.RegistrationHandler)

	e.GET("/api/register/confirmation", auth.RegistrationConfirmationHandler)

	e.POST("/api/test/match", match.FindMatchesHandlerTest)

	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	//e.Logger.Fatal(e.StartAutoTLS(":443"))
	e.Logger.Fatal(e.StartServer(s))
}
