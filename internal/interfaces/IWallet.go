package interfaces

import (
	"context"
	"ewallet-wallet/internal/models"

	"github.com/gin-gonic/gin"
)

type IWalletRepossitory interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
}

type IWalletService interface {
	Create(ctx context.Context, wallet *models.Wallet) error
}

type IWalletApi interface {
	Create(c *gin.Context)
}
