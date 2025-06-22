package services

import (
	"context"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/models"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type WalletService struct {
	WalletRepository interfaces.IWalletRepossitory
}

func (s *WalletService) Create(ctx context.Context, wallet *models.Wallet) error {
	return s.WalletRepository.CreateWallet(ctx, wallet)
}

func (s *WalletService) CreaditBalance(ctx context.Context, userID int, req models.TransactiontRequest) (models.BalanceResponse, error) {
	var response models.BalanceResponse

	TrxReference, err := s.WalletRepository.GetWalletTransactionByReference(ctx, req.Reference)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return response, errors.Wrap(err, "failed to get wallet transaction by reference")
		}
	}

	if TrxReference.ID > 0 {
		return response, errors.New("reference already exists, please use another reference")
	}
	wallet, err := s.WalletRepository.UpdateBalance(ctx, userID, req.Amount)
	if err != nil {
		return response, errors.Wrap(err, "failed to update wallet balance")
	}

	walletTansaction := &models.WalletTransaction{
		WalletID:              wallet.ID,
		Amount:                req.Amount,
		Reference:             req.Reference,
		WalletTransactionType: "CREDIT",
	}
	err = s.WalletRepository.CreateWalletTransaction(ctx, walletTansaction)
	if err != nil {
		return response, errors.Wrap(err, "failed to create wallet transaction")

	}
	response.Amount = req.Amount + wallet.Balance
	return response, nil
}
func (s *WalletService) DebitBalance(ctx context.Context, userID int, req models.TransactiontRequest) (models.BalanceResponse, error) {
	var response models.BalanceResponse

	TrxReference, err := s.WalletRepository.GetWalletTransactionByReference(ctx, req.Reference)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return response, errors.Wrap(err, "failed to get wallet transaction by reference")
		}
	}

	if TrxReference.ID > 0 {
		return response, errors.New("reference already exists, please use another reference")
	}
	wallet, err := s.WalletRepository.UpdateBalance(ctx, userID, -req.Amount)
	if err != nil {
		return response, errors.Wrap(err, "failed to update wallet balance")
	}

	walletTansaction := &models.WalletTransaction{
		WalletID:              wallet.ID,
		Amount:                req.Amount,
		Reference:             req.Reference,
		WalletTransactionType: "DEBIT",
	}
	err = s.WalletRepository.CreateWalletTransaction(ctx, walletTansaction)
	if err != nil {
		return response, errors.Wrap(err, "failed to create wallet transaction")

	}
	response.Amount = wallet.Balance - req.Amount
	return response, nil
}

func (s *WalletService) GetWalletBalance(ctx context.Context, userID int) (models.BalanceResponse, error) {
	var response models.BalanceResponse
	fmt.Println("User ID:", userID)
	wallet, err := s.WalletRepository.GetWalletBalanceByUserID(ctx, userID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return response, errors.Wrap(err, "failed to get wallet transaction by reference")
		}
	}

	response.Amount = wallet.Balance
	return response, nil
}
func (s *WalletService) GetWalletHistory(ctx context.Context, userID int, param models.WalletHistoryParam) ([]models.WalletTransaction, error) {
	var response []models.WalletTransaction

	wallet, err := s.WalletRepository.GetWalletBalanceByUserID(ctx, userID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return response, errors.Wrap(err, "failed to get wallet balance")
		}

	}

	offset := (param.Page - 1) * (param.Limit)
	response, err = s.WalletRepository.GetWalletHistory(ctx, wallet.ID, offset, param.Limit, param.WalletTransactionType)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return response, errors.Wrap(err, "failed to get wallet transaction history")
		}
	}
	return response, nil
}
