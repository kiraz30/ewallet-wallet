package models

import "time"

type Wallet struct {
	ID        int
	UserID    int     `gorm:"column:user_id`
	Balance   float64 `gorm:"column:balance;type:decimal(15,2)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (w *Wallet) TableName() string {
	return "wallets"
}

type WalletTransaction struct {
	ID                    int
	WalletID              int     `gorm:"column:wallet_id`
	Amount                float64 `gorm:"column:amount;type:decimal(15,2)"`
	WalletTransactionType string  `gorm:"column:wallet_transaction_type; type:ENUM('CREDIT', 'DEBIT')"`
	Reference             string  `gorm:"column:reference_id;type:varchar(100)"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (w *WalletTransaction) TableName() string {
	return "wallet_transactions"
}
