package entity

type Product struct {
	Id        string  `bson:"_id" json:"id"`
	Name      string  `bson:"name" json:"name"`
	UnitValue float64 `bson:"unitValue" json:"unitValue"`
}
