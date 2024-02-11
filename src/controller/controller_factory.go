package controller

import "github.com/geffersonFerraz/grinha-de-backend-2024-q1-demode/src/usecase"

type (
	Controller interface {
		NewTransactionController() TransactionsController
	}

	controller struct {
		useCase usecase.UseCases
	}
)

func NewController(uc usecase.UseCases) Controller {
	return &controller{useCase: uc}
}

func (c *controller) NewTransactionController() TransactionsController {
	return NewTransactionController(c.useCase.NewTransactionsUseCase())
}
