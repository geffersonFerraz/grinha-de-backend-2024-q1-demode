package usecase

import (
	"github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/repository"
)

type (
	UseCases interface {
		NewTransactionsUseCase() TransactionsUseCase
	}

	useCases struct {
		repos repository.Repositories
	}
)

func NewUseCase(repo repository.Repositories) UseCases {
	return &useCases{repos: repo}
}

// NewTransactionsRepository implements Repositories.
func (r *useCases) NewTransactionsUseCase() TransactionsUseCase {
	return NewTransactions(r.repos.NewTransactionsRepository())
}
