/*
Exercise link: https://docs.google.com/document/d/193B2OXkQIJ2Va1JjxZ9iVDYKvaF1TwRQcaDqkg-HMzA/edit#heading=h.tytft2pkn5l9
Problem 1:
*/

package main

import (
	"exercise1/Config"
	"exercise1/Models"
	"exercise1/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
