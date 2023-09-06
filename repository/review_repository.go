package repository

import (
	"context"
	"teduh-mongodb-assessment/contract"
	"teduh-mongodb-assessment/entities"
	"teduh-mongodb-assessment/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type reviewRepository struct {
	collection *mongo.Collection
}

func NewReviewRepository(db *mongo.Database) contract.ReviewRepository {
	return &reviewRepository{
		collection: db.Collection("reviews"),
	}
}

func (r *reviewRepository) FindAll(ctx context.Context) []entities.Review {
	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 404,
			Error:      err.Error(),
		})
	}

	var results []entities.Review

	if err := cursor.All(ctx, &results); err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 404,
			Error:      err.Error(),
		})
	}

	return results
}

func (r *reviewRepository) Insert(ctx context.Context, result entities.Review) {

	_, err := r.collection.InsertOne(ctx, result)

	if err != nil {

		panic(err)
	}

}
