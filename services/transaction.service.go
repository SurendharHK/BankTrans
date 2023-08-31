package services

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/interfaces"
	"main.go/models"
)

type TransactionService struct {
	client                *mongo.Client
	CustomerCollection    *mongo.Collection
	TransactionCollection *mongo.Collection
	ctx                   context.Context
}

func NewTransactionServiceInit(client *mongo.Client, Customercollection *mongo.Collection, TransactionCollection *mongo.Collection, ctx context.Context) interfaces.ITransaction {
	return &TransactionService{
		client:                client,
		CustomerCollection:    Customercollection,
		TransactionCollection: TransactionCollection,
		ctx:                   ctx,
	}

}

// Transfer implements interfaces.IBank.
func (a *TransactionService) Transfer(from string, to string, amount float64) (string, error) {

	session, err := a.client.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(context.Background())

	filter := bson.M{"cus_id": from}

   var account *models.Account

    err1 := a.CustomerCollection.FindOne(context.Background(), filter).Decode(&account)
    if err1 != nil {
        return "error", err1
    }
	
	if account.Balance > amount{


	_, err = session.WithTransaction(context.Background(), func(ctx mongo.SessionContext) (interface{}, error) {
		_, err := a.CustomerCollection.UpdateOne(context.Background(),
			bson.M{"cus_id": from},
			bson.M{"$inc": bson.M{"balance": -amount}})
		if err != nil {
			fmt.Println("Transaction Failed")
			return nil, err
		}
		_, err2 := a.CustomerCollection.UpdateOne(context.Background(), bson.M{"cus_id": to}, bson.M{"$inc": bson.M{"balance": amount}})

		if err2 != nil {
			fmt.Println("Transaction Failed")
			return nil, err2
		}
		transactionToInsert := models.Transaction{
			ID:     "T0011",
			From: from,
			To:   to,
			Amount: amount,
		}
		res, _ := a.TransactionCollection.InsertOne(context.Background(), &transactionToInsert)
		var newUser *models.Transaction
		query := bson.M{"_id": res.InsertedID}

		err3 := a.TransactionCollection.FindOne(a.ctx, query).Decode(&newUser)
		if err3 != nil {
			return nil, err3
		}
		return newUser, nil
	})
	if err != nil {
		return "failed", err
	}}else{
		return "Insufficent balance",nil
	}
	return "success", nil
}
