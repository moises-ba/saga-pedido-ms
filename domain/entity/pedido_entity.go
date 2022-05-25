package entity

import "time"

type Pedido struct {
	Id            string         `bson:"_id" json:"id"`
	User          *User          `bson:"user" json:"user"`
	Items         []*Item        `bson:"items" json:"items"`
	Status        string         `bson:"status" json:"status"`
	Reason        string         `json:"reason"`
	Date          time.Time      `bson:"date" json:"date"`
	TotalPrice    float64        `bson:"totalPrice" json:"totalPrice"`
	PaymentDetail *PaymentDetail `bson:"paymentDetail" json:"paymentDetail"`
}
