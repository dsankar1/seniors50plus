package main

import (
	"fmt"
	"seniors50plus/internal/models"
)

func main() {
	email := "daryan.sankar1@gmail.com"
	fmt.Println(email)
	dbc := models.NewDatabaseConnection()
	fmt.Println(dbc.CreateTable(models.RoommateOffer{}, models.CommunicationRequest{}, models.Resident{}))
}
