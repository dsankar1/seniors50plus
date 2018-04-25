package main

import (
	"fmt"
	"seniors50plus/internal/models"
	"time"
)

func main() {
	offer := models.RoommateOffer{
		ID:                5,
		UploaderID:        1,
		State:             "GA",
		City:              "Alpharetta",
		Zip:               30005,
		GenderRequirement: "male",
		PropertyType:      "apartment",
		CreatedAt:         time.Now(),
	}
	fmt.Println(offer)
	elastic := models.NewElasticClient()
	//elastic.Delete("4")
	results := models.QueryResponse{}
	if err := elastic.Get(&offer, &results); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(results)
	/*err := elastic.Put(&offer)
	if err != nil {
		fmt.Println(err)
	}*/
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
