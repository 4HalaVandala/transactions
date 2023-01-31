package service

import (
	"github.com/4halavandala/transactions/tx-app/domain/model"
	"github.com/4halavandala/transactions/tx-app/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(transaction model.Transaction) (int, error) {
	return s.repo.Create(transaction)
}

func (s *TransactionService) GetAll() ([]model.Transaction, error) {
	return s.repo.GetAll()
}

func (s *TransactionService) GetById(id int) (model.Transaction, error) {
	return s.repo.GetById(id)
}
