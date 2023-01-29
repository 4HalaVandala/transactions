package model

import "time"

type TransactionType int

const (
	Deposit TransactionType = iota + 1
	Withdrawal
)

type Transaction struct {
	Id        int             `gorm:"primary key" json:"id" db:"id"`
	TxType    TransactionType `json:"tx_type" binding:"required"`
	Amount    int             `json:"amount" binding:"required"`
	CreatedAt time.Time       `json:"created_at" binding:"required"`
	UpdatedAt time.Time       `json:"updated_at" binding:"required"`
}

/* Transaction types :
1 - deposit
2 - withdrawal
*/
