package entity

type Item struct {
	Product    Product `json:"product"`
	Quantity   int32   `json:"quantity"`
	TotalPrice float64 `json:"totalPrice"`
}
