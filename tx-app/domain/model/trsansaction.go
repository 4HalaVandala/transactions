package model

import "time"

type Transaction struct {
	Id         int64     `gorm:"primary key" json:"id" db:"id"`
	Username   string    `json:"username" binding:"required"`
	Amount     int64     `json:"amount,string,omitempty" binding:"required"`
	TxType_id  int64     `json:"txtype_id,string,omitempty" binging:"required"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

/* Transaction types :
1 - deposit
2 - withdrawal
*/

type transaction_type struct {
	Id   int64  `gorm:"primary key" json:"id" db:"id"`
	Type string `json:"transactionType"`
}
