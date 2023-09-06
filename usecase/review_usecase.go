package usecase

import (
	"context"
	"teduh-mongodb-assessment/contract"
	"teduh-mongodb-assessment/entities"
	"teduh-mongodb-assessment/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type reviewUseCase struct {
	repo contract.ReviewRepository
}

func NewReviewUseCase(s *contract.ReviewRepository) contract.ReviewUseCase {

	return &reviewUseCase{
		repo: *s,
	}
}

func (r *reviewUseCase) List() []entities.Review {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)

	defer cancel()
	results := r.repo.FindAll(ctx)

	return results
}

func (r *reviewUseCase) Create(payload model.CreateReviewModel) {

	id, err := primitive.ObjectIDFromHex(payload.PsiId)

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 400,
			Error:      "Invalid psikolog id",
		})
	}

	var data entities.Review = entities.Review{
		PsiId:   id,
		Rating:  payload.Rating,
		Message: payload.Message,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	r.repo.Insert(ctx, data)
}
