package main

import (
	"fmt"
	"seniors50plus/internal/models"
)

func main() {
	email := "daryan.sankar1@gmail.com"
	if db, err := models.NewDatabaseConnection(); err != nil {
		fmt.Print(err)
	} else {
		defer db.Close()
		//db.CreateTable(&models.User{})
		//db.CreateTable(&models.Tag{})
		var user models.User
		user = models.User{
			Email: email,
			Tags: []models.Tag{
				{
					Content: "Test 1",
				},
				{
					Content: "Test 2",
				},
				{
					Content: "Test 3",
				},
			},
		}
		db.Create(&user)
		//var tags []models.Tag
		//db.Where("owner = ?", email).Find(&tags)
		//fmt.Println(tags)
		//var tags models.User
		//db.Where("email = ?", email).First(&tags)
		//db.Where("email = ?", email).First(&user)

		//fmt.Println(tags)
	}

}
