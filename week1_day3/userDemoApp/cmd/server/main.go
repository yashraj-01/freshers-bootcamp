package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"yash/userDemoApp/config"
	"yash/userDemoApp/internal/core/listeners/http"
	httpTransport "yash/userDemoApp/internal/transport/http"
	"yash/userDemoApp/internal/users"
	"yash/userDemoApp/internal/users/store"
)

func main() {
	fmt.Println("Running")
	initDatabase()
}

func initDatabase() {
	db, err := config.New()
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer db.Close()

	usersStore := store.New(db)
	usersInstance := users.New(usersStore)

	httpServer := httpTransport.New(usersInstance)

	err = http.New(httpServer)

	if err != nil {
		fmt.Println("Status:", err)
	}
}
