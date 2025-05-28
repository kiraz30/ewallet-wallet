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
}

type IWalletService interface {
	Create(ctx context.Context, wallet *models.Wallet) error
	CreaditBalance(ctx context.Context, userID int, req models.TransactiontRequest) (models.TransactiontResponse, error)
	DebitBalance(ctx context.Context, userID int, req models.TransactiontRequest) (models.TransactiontResponse, error)
}

type IWalletApi interface {
	Create(c *gin.Context)
	CreaditBalance(c *gin.Context)
	DebitBalance(c *gin.Context)
}
