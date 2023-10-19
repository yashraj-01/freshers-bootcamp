package main

import (
	"fmt"
	"yash/shopApp/config"
)

func main() {
	fmt.Println("Running...")
	initDatabase()
}

func initDatabase() {
	db, err := config.New()

	if err != nil {
		fmt.Println("Status:", err)
	} else {
		fmt.Println("Connected to database")
	}

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

}
