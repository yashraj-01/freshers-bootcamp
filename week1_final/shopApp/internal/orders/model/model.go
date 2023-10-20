package model

type Order struct {
	Id         uint   `json:"id"`
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
	Quantity   string `json:"quantity"`
	Status     string `json:"status"`
}
