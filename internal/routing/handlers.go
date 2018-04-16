package routing

import (
	"seniors50plus/internal/auth"
	"seniors50plus/internal/offer"
	"seniors50plus/internal/request"
	"seniors50plus/internal/user"

	"github.com/labstack/echo"
)

func RegisterHandlers(e *echo.Echo) {

	// AUTH
	e.POST("/api/auth/login", auth.LoginHandler)

	e.POST("/api/auth/signup", auth.SignupHandler)

	e.GET("/api/auth/signup/confirmation", auth.EmailConfirmationHandler)

	// USER
	e.GET("/api/user/:id", user.GetUserHandler)

	e.POST("/api/user/list", user.GetUserListHandler)

	e.PUT("/api/user", user.UpdateUserHandler)

	// OFFERS
	e.POST("/api/offer", offer.PostOfferHandler)

	e.GET("/api/offer", offer.GetOfferHandler)

	e.GET("/api/offer/:id", offer.GetOfferHandler)

	e.DELETE("/api/offer", offer.DeleteOfferHandler)

	// COMMUNICATION REQUESTS
	e.POST("/api/offer/:id/request", request.CreateCommunicationRequestHandler)

	e.DELETE("/api/offer/:id/request", request.DeleteCommunicationRequestHandler)

	e.PUT("/api/offer/request/:id", request.RespondToCommunicationRequestHandler) //?status=value

	// RESIDENT REQUESTS
	e.POST("/api/user/:id/request", request.CreateResidentRequestHandler)

	e.DELETE("/api/user/:id/request", request.DeleteResidentRequestHandler)

	e.PUT("/api/user/request/:id", request.RespondToResidentRequestHandler) //?status=value

}
