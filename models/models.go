package models

type Order struct {
	ID           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	Items        []Item `json:"items"`
}

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
}
