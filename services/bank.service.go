package services

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/interfaces"
	"main.go/models"
)

type BankingService struct {
	BankingCollection *mongo.Collection
	ctx               context.Context
}



// GetTransactionByDate implements interfaces.IBank.

// UpdateManyCustomer implements interfaces.IBank.
func (b *BankingService) InsertManyCustomer(customers []interface{}) (*mongo.InsertManyResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := b.BankingCollection.InsertMany(ctx, customers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil

}

// DeleteCustomer implements interfaces.IBank.
func (t *BankingService) DeleteCustomer(name string) (*mongo.DeleteResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"name", name}}

	result, err := t.BankingCollection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result, nil

}

// UpdateCustomer implements interfaces.IBank.
func (t *BankingService) UpdateCustomer(intialName string, updateName string) (*mongo.UpdateResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"name", intialName}}
	update := bson.D{{"$set", bson.D{{"name", updateName}}}}
	result, err := t.BankingCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result, nil

}

// GetCustomers implements interfaces.IBank.
func (t *BankingService) GetCustomers() ([]*models.Account, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{}}
	result, err := t.BankingCollection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		//do something
		fmt.Println(result)
		//build the array of products for the cursor that we received.
		var customers []*models.Account
		for result.Next(ctx) {
			product := &models.Account{}
			err := result.Decode(product)

			if err != nil {
				return nil, err
			}
			//fmt.Println(product)
			customers = append(customers, product)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(customers) == 0 {
			return []*models.Account{}, nil
		}

		return customers, nil
	}
}

func NewBankingServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.IBank {
	return &BankingService{collection, ctx}

}

// CreateTransaction implements interfaces.ITransaction.
func (t *BankingService) CreateCustomer(customer *models.Account) (string, error) {

	//hashedPassword, _ := utils.HashPassword(user.Password)

	_, err := t.BankingCollection.InsertOne(t.ctx, &customer)
	if err != nil {
		return "error", nil
	}

	return "success", nil
}
