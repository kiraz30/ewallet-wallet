package repository

import (
	"context"
	"ewallet-wallet/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type WalletRepository struct {
	DB *gorm.DB
}

func (r *WalletRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	return r.DB.Create(wallet).Error
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, userID int, amount float64) (models.Wallet, error) {
	var wallet models.Wallet
	err := r.DB.Transaction(func(tx *gorm.DB) error {

		err := tx.Raw("SELECT id, user_id, balance FROM wallets WHERE user_id = ? FOR UPDATE", userID).Scan(&wallet).Error
		if err != nil {
			return err
		}

		if (wallet.Balance + amount) < 0 {
			return fmt.Errorf("current balance is not enough : %f - %f", wallet.Balance, amount)
		}
		err = tx.Exec("UPDATE wallets SET balance = balance + ? WHERE user_id = ?", amount, userID).Error
		if err != nil {
			return err
		}
		return nil
	})
	return wallet, err
}
func (r *WalletRepository) CreateWalletTransaction(ctx context.Context, walletTransaction *models.WalletTransaction) error {
	return r.DB.Create(walletTransaction).Error
}

func (r *WalletRepository) GetWalletTransactionByReference(ctx context.Context, reference string) (models.WalletTransaction, error) {
	var WalletTransaction models.WalletTransaction
	err := r.DB.Where("reference_id = ?", reference).Last(&WalletTransaction).Error
	return WalletTransaction, err
}

func (r *WalletRepository) GetWalletBalanceByUserID(ctx context.Context, userID int) (models.Wallet, error) {
	var wallet models.Wallet
	fmt.Println("UserID:", userID)
	err := r.DB.Where("user_id = ?", userID).Last(&wallet).Error
	return wallet, err
}

func (r *WalletRepository) GetWalletHistory(ctx context.Context, walletID int, offset, limit int, transactionType string) ([]models.WalletTransaction, error) {
	var walletHistory []models.WalletTransaction
	sql := r.DB
	if transactionType != "" {
		sql = sql.Where("wallet_transaction_type= ?", transactionType)
	}
	fmt.Println("Limit:", limit)
	err := sql.Debug().Limit(limit).Offset(offset).Order("id DESC").Find(&walletHistory).Error
	return walletHistory, err
}
