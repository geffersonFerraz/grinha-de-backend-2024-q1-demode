package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Repositories interface {
		NewTransactionsRepository() TransactionsRepository
	}

	repositories struct {
		DB *mongo.Client
	}
)

func NewRepositories(db *mongo.Client) Repositories {
	return &repositories{DB: db}
}

// NewTransactionsRepository implements Repositories.
func (r *repositories) NewTransactionsRepository() TransactionsRepository {
	return NewTransactions(r.DB)
}
