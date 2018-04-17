package main

import (
	"fmt"
	"seniors50plus/internal/models"
)

func main() {
	email := "daryan.sankar1@gmail.com"
	fmt.Println(email)
	dbc := models.NewDatabaseConnection()
	if err := dbc.Open(); err != nil {
		fmt.Println(err.Error())
	}
	defer dbc.Close()
	fmt.Println(dbc.CreateTable(models.Report{}, models.Ban{}, models.Flag{}))
}
