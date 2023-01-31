package service

import (
	"github.com/4halavandala/transactions/tx-app/domain/model"
	"github.com/4halavandala/transactions/tx-app/pkg/repository"
)

type Transaction interface {
	Create(transaction model.Transaction) (int, error)
	GetById(id int) (model.Transaction, error)
	GetAll() ([]model.Transaction, error)
}

type Service struct {
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Transaction: NewTransactionService(repos.Transaction),
	}
}
