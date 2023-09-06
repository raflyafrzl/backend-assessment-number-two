package repository

import (
	"context"
	"teduh-mongodb-assessment/contract"
	"teduh-mongodb-assessment/entities"
	"teduh-mongodb-assessment/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type psikologRepository struct {
	collection *mongo.Collection
}

func NewPsikologRepostory(d *mongo.Database) contract.PsikologRepository {

	return &psikologRepository{
		collection: d.Collection("psikolog"),
	}

}

func (a *psikologRepository) Insert(ctx context.Context, name string) error {

	var data entities.Psikolog = entities.Psikolog{
		Name: name,
	}

	_, err := a.collection.InsertOne(context.TODO(), data)

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 400,
			Error:      err.Error(),
		})
	}

	return nil

}

func (a *psikologRepository) FindAll(ctx context.Context) []bson.M {

	var pipeline []bson.M = []bson.M{
		{
			"$lookup": bson.M{
				"from":         "reviews",
				"localField":   "_id",
				"foreignField": "psikolog_id",
				"as":           "reviews",
			},
		},
	}

	// var pipeline []primitive.M = []bson.M{
	// 	{
	// 		"$lookup": bson.M{
	// 			"from":         "reviews",
	// 			"localField":   "_id",
	// 			"foreignField": "psikolog_id",
	// 			"as":           "reviews",
	// 		},
	// 	},
	// }

	cursor, err := a.collection.Aggregate(ctx, pipeline)

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 404,
			Error:      err.Error(),
		})
	}

	var docs []bson.M
	if err := cursor.All(ctx, &docs); err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 404,
			Error:      err.Error(),
		})
	}

	return docs
}
