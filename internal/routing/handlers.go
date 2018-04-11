package routing

import (
	"seniors50plus/internal/auth"
	"seniors50plus/internal/user"

	"github.com/labstack/echo"
)

func RegisterHandlers(e *echo.Echo) {

	e.POST("/api/auth/login", auth.LoginHandler)

	e.POST("/api/auth/signup", auth.SignupHandler)

	e.GET("/api/auth/confirmation", auth.EmailConfirmationHandler)

	e.PUT("/api/user/modify", user.ModifyUserHandler)

}
