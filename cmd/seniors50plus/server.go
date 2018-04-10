package main

import (
	"net/http"
	"time"

	"seniors50plus/internal/routing"

	"github.com/labstack/echo"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("roommates40plus.com")
	e.AutoTLSManager.Cache = autocert.DirCache("./var/www/.cache")

	routing.RegisterMiddleware(e)
	routing.RegisterHandlers(e)

	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
	//e.Logger.Fatal(e.StartAutoTLS(":443"))
}
