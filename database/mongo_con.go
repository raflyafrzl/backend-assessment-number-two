package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()
	option := options.Client().ApplyURI(os.Getenv("MONGO_ADDR"))

	client, err := mongo.Connect(ctx, option)

	if err != nil {
		panic(err)
	}

	var db *mongo.Database = client.Database(os.Getenv("MONGO_DB"))

	// cmd := bson.D{{"create", "psikolog"}}

	// var result bson.M

	return db

}
