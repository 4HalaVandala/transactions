package service

import (
	"github.com/4halavandala/transactions/domain/model"
	"github.com/4halavandala/transactions/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(id int, transaction model.Transaction) (int, error) {
	return s.repo.Create(id, transaction)
}

func (s *TransactionService) GetAll() ([]model.Transaction, error) {
	return s.repo.GetAll()
}

func (s *TransactionService) GetById(id int) (model.Transaction, error) {
	return s.repo.GetById(id)
}
