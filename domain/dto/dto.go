package dto

import "time"

type Product struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	UnitPrice float64 `json:"unitPrice"`
}

type Item struct {
	Product   *Product `json:"product"`
	Quantity  int32    `json:"quantity"`
	UnitPrice float64  `json:"unitPrice"`
	Total     float64  `json:"total"` //preco calculado Quantity * UnitPrice
}

type PaymentDetail struct {
	CardType   string `json:"cardType"`   //VISA etc
	CardNumber string `json:"cardNumber"` //000000
}

type User struct {
	Id string `json:"id"`
}

type PedidoRequest struct {
	Id            string         `json:"id"`
	User          *User          `json:"user"`
	Items         []*Item        `json:"items"`
	Status        string         `json:"status"`
	Reason        string         `json:"reason"`
	Date          time.Time      `json:"date"`
	PaymentDetail *PaymentDetail `json:"paymentDetail"`
	TotalPrice    float64        `json:"totalPrice"`
}

type PedidoResponse struct {
	Id            string         `json:"id"`
	User          *User          `json:"user"`
	Items         []*Item        `json:"items"`
	Status        string         `json:"status"`
	Reason        string         `json:"reason"`
	Date          time.Time      `json:"date"`
	PaymentDetail *PaymentDetail `json:"paymentDetail"`
	TotalPrice    float64        `json:"totalPrice"`
}

type PedidoCreatedResponse struct {
	Id string `json:"id"`
}
