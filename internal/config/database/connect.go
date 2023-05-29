package database

import (
	"context"
	"ecommerce/internal/config"
	"fmt"
	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ecommerceDB *mongo.Database

// ConnectMongoDBEcommerce ...
func ConnectMongoDBEcommerce() {
	cfg := config.GetENV().MongoDB
	connectOptions := options.ClientOptions{}
	client, err := mongo.Connect(context.Background(), connectOptions.ApplyURI(cfg.URI))
	if err != nil {
		fmt.Println("Error when connect to MongoDB database", cfg.URI, err)
		panic(err)
	}

	fmt.Println(aurora.Green("*** CONNECTED TO MONGODB: " + cfg.URI + " --- DB: " + cfg.DBName))
	ecommerceDB = client.Database(cfg.DBName)
	index()
}

func GetMongoDBEcommerce() *mongo.Database {
	return ecommerceDB
}

func GetUserCol() *mongo.Collection {
	return ecommerceDB.Collection(colUser)
}
