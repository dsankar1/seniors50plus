package main

import (
	"fmt"
	"net/http"
	"seniors50plus/internal/models"
	"time"

	"github.com/labstack/echo"
)

func main() {
	email := "daryan.sankar1@gmail.com"
	fmt.Println(email)
	dbc := models.NewDatabaseConnection()
	offer := models.RoommateOffer{
		TargetOccupantCount: 1,
	}
	offers, err := dbc.QueryOffers(&offer)
	if err != nil {
		fmt.Println(err)
	}
	/*if err := dbc.AttachResidentsOccupants(&offer); err != nil {
		fmt.Println(err)
	}*/

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, offers)
	})
	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
	/*dbstring := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v", "capstoneuser",
		os.Getenv("DB_PASSWORD"),
		"capstone.cczajq2nppkf.us-east-2.rds.amazonaws.com",
		"roommates40plusv2?charset=utf8&parseTime=true",
	)
	if db, err := gorm.Open("mysql", dbstring); err != nil {
		fmt.Println(err.Error())
	} else {
		defer db.Close()
		//db.CreateTable(&models.Resident{})
		//db.CreateTable(&models.RoommateOffer{})
		offer := models.RoommateOffer{
			PosterEmail: email,
		}
		db.Create(&offer)
		fmt.Println(offer.Residents[0].Occupant)
	}*/
}
