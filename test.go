package main

import (
	"fmt"
	"seniors50plus/internal/models"
)

func main() {
	offer := models.RoommateOffer{
		ID:         1,
		UploaderID: 1,
		State:      "GA",
		City:       "Alpharetta",
		Zip:        30005,
	}
	elastic := models.NewElasticClient()

	if res, err := elastic.Get(&offer); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res.Hits)
	}
	//res, _ := elastic.Put(&offer)
	//fmt.Println(res)
	/*if err := elastic.Put(&offer); err != nil {
		fmt.Println(err.Error())
	}
	/*email := "daryan.sankar1@gmail.com"
	fmt.Println(email)
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		fmt.Println(err.Error())
	}
	defer dbc.Close()
	fmt.Println(dbc.CreateTable(models.RoommateOffer{}))*/
}
