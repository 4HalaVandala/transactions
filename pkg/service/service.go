package service

import (
	"github.com/4halavandala/transactions/domain/model"
	"github.com/4halavandala/transactions/pkg/repository"
)

type Transaction interface {
	Create(id int, transaction model.Transaction) (int, error)
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
