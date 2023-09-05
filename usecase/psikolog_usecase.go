package usecase

import (
	"context"
	"teduh-mongodb-assessment/contract"
	"teduh-mongodb-assessment/entities"
	"time"
)

type psikologUseCase struct {
	repository contract.PsikologRepository
}

func NewPsikologUseCase(r *contract.PsikologRepository) contract.PsikologService {

	return &psikologUseCase{
		repository: *r,
	}

}

func (a *psikologUseCase) Create(name string) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	a.repository.Insert(ctx, name)

}

func (a *psikologUseCase) List() []entities.Psikolog {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()
	return a.repository.FindAll(ctx)

}
