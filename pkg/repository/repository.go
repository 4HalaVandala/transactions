package repository

import (
	"github.com/4halavandala/transactions/domain/model"
	"github.com/jmoiron/sqlx"
)

type Transaction interface {
	Create(id int, transaction model.Transaction) (int, error)
	GetAll() ([]model.Transaction, error)
	GetById(id int) (model.Transaction, error)
}

type Repository struct {
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransactionPostgres(db),
	}
}
