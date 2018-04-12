package main

import (
	"fmt"
	"seniors50plus/internal/models"
)

func main() {
	if db, err := models.NewDatabaseConnection(); err != nil {
		fmt.Print(err)
	} else {
		defer db.Close()
		/*user := Test{
			Email:     "daryan.sankar1@gmail.com",
			Firstname: "Daryan",
			Lastname:  "Sankar",
			Birthdate: "1932-02-02",
			Gender:    "male",
		}*/
		db.CreateTable(&Test{})
		//results := db.Create(&user)
		//results := db.First(&user)
		//fmt.Println(results)
	}

}

type Test struct {
	Email     string
	Firstname string
	Lastname  string
	Birthdate string
	Gender    string
	Password  string `gorm:"default:''"`
}
