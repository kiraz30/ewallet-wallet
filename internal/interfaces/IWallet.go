package interfaces

import (
	"context"
	"ewallet-wallet/internal/models"

	"github.com/gin-gonic/gin"
)

type IWalletRepossitory interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	UpdateBalance(ctx context.Context, userID int, amount float64) (models.Wallet, error)
	CreateWalletTransaction(ctx context.Context, walletTransaction *models.WalletTransaction) error
	GetWalletTransactionByReference(ctx context.Context, reference string) (models.WalletTransaction, error)
	GetWalletBalanceByUserID(ctx context.Context, userID int) (models.Wallet, error)
	GetWalletHistory(ctx context.Context, walletID int, offset, limit int, transactionType string) ([]models.WalletTransaction, error)
}

type IWalletService interface {
	Create(ctx context.Context, wallet *models.Wallet) error
	CreaditBalance(ctx context.Context, userID int, req models.TransactiontRequest) (models.BalanceResponse, error)
	DebitBalance(ctx context.Context, userID int, req models.TransactiontRequest) (models.BalanceResponse, error)
	GetWalletBalance(ctx context.Context, userID int) (models.BalanceResponse, error)
	GetWalletHistory(ctx context.Context, userID int, param models.WalletHistoryParam) ([]models.WalletTransaction, error)
}

type IWalletApi interface {
	Create(c *gin.Context)
	CreaditBalance(c *gin.Context)
	DebitBalance(c *gin.Context)
	GetWalletBalance(c *gin.Context)
	GetWalletHistory(c *gin.Context)
}
