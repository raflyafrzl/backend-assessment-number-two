package contract

import (
	"context"
	"teduh-mongodb-assessment/entities"
)

type PsikologService interface {
	List() []entities.Psikolog
	Create(name string)
}

type PsikologRepository interface {
	FindAll(ctx context.Context) []entities.Psikolog
	Insert(ctx context.Context, name string) error
}
