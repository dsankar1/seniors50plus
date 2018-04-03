package match

import (
	"fmt"
	"net/http"
	"seniors50plus/internal/auth"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func FindMatchesHandlerTest(c echo.Context) error {
	r := new(MatchRequest)
	if err := c.Bind(r); err != nil {
		return err
	}
	fmt.Println(r)
	if token, ok := c.Get("user").(*jwt.Token); ok {
		fmt.Println(token)
		var matches [10]Match
		for i := 0; i < cap(matches); i++ {
			offer := ExampleOffer
			offer.Id = uint64(i)
			offer.PostedBy = "MatchedUser" + strconv.Itoa(i) + "@gmail.com"
			offer.Occupants = []string{offer.PostedBy}
			match := Match{
				Offer:         offer,
				Compatability: 65.5 + float32(i),
			}
			matches[i] = match
		}
		return c.JSON(http.StatusOK, matches)
	} else {
		return c.JSON(http.StatusBadRequest, auth.Error{Message: "Type assertion for JWT failed"})
	}
}

func validateRequest(r *MatchRequest) bool {
	return true
}
