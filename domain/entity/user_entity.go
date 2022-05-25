package entity

type User struct {
	Id    string `bson:"_id" json:"id"`
	Email string `bson:"email" json:"email"`
	Login string `bson:"login" json:"login"`
}
