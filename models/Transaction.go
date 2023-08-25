package models

import "time"

type Transaction struct {
	ID        string    `json:"id" bson:"id"`
	From      string    `json:"from" bson:"from"`
	To        string    `json:"to" bson:"to"`
	Amount    float64   `json:"amount" bson:"amount"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

type Paying struct{
	From      string    `json:"from" bson:"from"`
	To        string    `json:"to" bson:"to"`
	Amount    float64   `json:"amount" bson:"amount"`
}
