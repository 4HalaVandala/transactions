package repository

import (
	"fmt"
	"time"

	"github.com/4halavandala/transactions/tx-app/domain/model"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) Create(transaction model.Transaction) (int, error) {
	var id int

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createTxQuery := fmt.Sprintf("INSERT INTO %s (username, txType_id, amount, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id", transactionsTable)
	row := tx.QueryRow(
		createTxQuery,
		transaction.Username,
		transaction.TxType_id,
		transaction.Amount,
		time.Now(),
		time.Now())

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TransactionPostgres) GetAll() ([]model.Transaction, error) {
	var txList []model.Transaction

	query := fmt.Sprintf("SELECT * FROM %s", transactionsTable)

	err := r.db.Select(&txList, query)

	return txList, err
}

func (r *TransactionPostgres) GetById(id int) (model.Transaction, error) {
	var tx model.Transaction

	query := fmt.Sprintf("SELECT * FROM %s t WHERE t.id =$1",
		transactionsTable)

	err := r.db.Get(&tx, query, id)

	return tx, err
}
