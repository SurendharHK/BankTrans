package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/models"
)

type IBank interface {
	CreateCustomer(customer *models.Account) (string, error)
	GetCustomers() ([]*models.Account, error)
	UpdateCustomer(intialName string, updateName string) (*mongo.UpdateResult, error)
	DeleteCustomer(name string) (*mongo.DeleteResult, error)
	InsertManyCustomer(customers []interface{}) (*mongo.InsertManyResult, error)
	
}
