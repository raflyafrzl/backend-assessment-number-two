package contract

import (
	"context"
	"teduh-mongodb-assessment/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type PsikologUseCase interface {
	List() []entities.PsikologReview
	Create(name string)
}

type PsikologRepository interface {
	FindAll(ctx context.Context) []bson.M
	Insert(ctx context.Context, name string) error
}
