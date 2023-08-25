package models

type Account struct {
	ID      string  `json:"cus_id" bson:"cus_id"`
	Balance float64 `json:"balance" bson:"balance"`
}

type UpdateName struct {
	IntialName string `json:"intialname" bson:"intialname"`
	UpdateName string `json:"updatename" bson:"updatename"`
}
