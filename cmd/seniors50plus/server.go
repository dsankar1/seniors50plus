package main

import (
	"net/http"
	"seniors50plus/internal/auth"
	"seniors50plus/internal/middleware"

	"github.com/labstack/echo"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("roommates40plus.com")
	e.AutoTLSManager.Cache = autocert.DirCache("./var/www/.cache")
	middleware.ApplyMiddleware(e)

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to the api")
	})
	e.GET("/api/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.Path())
	})
	e.POST("/api/authenticate", auth.AuthenticationHandler)
	e.POST("/api/register", auth.RegistrationHandler)

	/*s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}*/
	e.Logger.Fatal(e.StartAutoTLS(":443"))
	//e.Logger.Fatal(e.StartServer(s))
}
