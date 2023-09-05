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

	_, err := a.collection.InsertOne(ctx, data)

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 400,
			Error:      err.Error(),
		})
	}

	return nil

}

func (a *psikologRepository) FindAll(ctx context.Context) []entities.Psikolog {

	cursor, err := a.collection.Find(ctx, bson.M{})

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 404,
			Error:      err.Error(),
		})
	}

	var results []entities.Psikolog

	if err := cursor.All(ctx, &results); err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 404,
			Error:      err.Error(),
		})
	}

	return results
}
