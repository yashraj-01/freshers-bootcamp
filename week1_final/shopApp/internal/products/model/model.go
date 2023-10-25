package model

type Product struct {
	Id          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Quantity    string `json:"quantity"`
}
