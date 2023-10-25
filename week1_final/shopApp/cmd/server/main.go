package main

import (
	"fmt"
	"yash/shopApp/config"
	"yash/shopApp/internal/core/listeners/http"
	"yash/shopApp/internal/customers"
	cStore "yash/shopApp/internal/customers/store"
	"yash/shopApp/internal/orders"
	oStore "yash/shopApp/internal/orders/store"
	"yash/shopApp/internal/products"
	pStore "yash/shopApp/internal/products/store"
	httpTransport "yash/shopApp/internal/transport/http"
)

func main() {
	fmt.Println("Running...")
	setupDatabase()
}

func setupDatabase() {
	db, err := config.New()

	if err != nil {
		fmt.Println("Status:", err)
	} else {
		fmt.Println("Connected to database")
	}

	defer func() {
		dbInstance, err := db.DB()
		if err != nil {
			fmt.Println("Failed to get database instance: ", err)
		} else {
			err = dbInstance.Close()
			if err != nil {
				fmt.Println("Failed to close the database: ", err)
			}
		}
	}()

	customersStore := cStore.New(db)
	customersInstance := customers.New(customersStore)

	ordersStore := oStore.New(db)
	ordersInstance := orders.New(ordersStore)

	productsStore := pStore.New(db)
	productsInstance := products.New(productsStore)

	httpServer := httpTransport.New(customersInstance, ordersInstance, productsInstance)

	err = http.New(httpServer)

	if err != nil {
		fmt.Println("Status:", err)
	}
}