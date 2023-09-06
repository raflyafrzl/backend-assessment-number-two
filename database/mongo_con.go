package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() *mongo.Database {

	option := options.Client().ApplyURI(os.Getenv("MONGO_ADDR"))

	client, err := mongo.Connect(context.TODO(), option)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	var db *mongo.Database = client.Database(os.Getenv("MONGO_DB"))

	// cmd := bson.D{{"create", "psikolog"}}

	// var result bson.M

	return db

}
