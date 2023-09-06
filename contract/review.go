package contract

import (
	"context"
	"teduh-mongodb-assessment/entities"
	"teduh-mongodb-assessment/model"
)

type ReviewUseCase interface {
	List() []entities.Review
	Create(model.CreateReviewModel)
}

type ReviewRepository interface {
	FindAll(ctx context.Context) []entities.Review
	Insert(context.Context, entities.Review)
}
