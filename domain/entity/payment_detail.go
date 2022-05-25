package entity

type PaymentDetail struct {
	CardType   string `bson:"cardType" json:"cardType"`     //VISA etc
	CardNumber string `bson:"cardNumber" json:"cardNumber"` //000000
}
