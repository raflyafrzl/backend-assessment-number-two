package usecase

import (
	"context"
	"teduh-mongodb-assessment/contract"
	"teduh-mongodb-assessment/entities"
	"time"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type psikologUseCase struct {
	repository contract.PsikologRepository
}

func NewPsikologUseCase(r *contract.PsikologRepository) contract.PsikologUseCase {

	return &psikologUseCase{
		repository: *r,
	}

}

func (a *psikologUseCase) Create(name string) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	a.repository.Insert(ctx, name)

}

func (a *psikologUseCase) List() []entities.PsikologReview {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	var docs []bson.M = a.repository.FindAll(ctx)

	var results []entities.PsikologReview
	var avg float32
	for index, data := range docs {
		mongoId := data["_id"].(primitive.ObjectID)
		results = append(results, entities.PsikologReview{
			Id:   mongoId.Hex(),
			Name: data["name"].(string),
		})

		mapstructure.Decode(data["reviews"], &results[index].Review)

		for _, d := range results[index].Review {
			avg = avg + float32(d["rating"].(int32))
		}

		avg = avg / float32(len(results[index].Review))
		results[index].Average = avg
	}

	return results

}
