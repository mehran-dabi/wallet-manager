package service

import (
	"context"
	"wallet-manager/domain/constants"
	"wallet-manager/domain/wallet/dto"
	"wallet-manager/domain/wallet/repository"
	"wallet-manager/domain/wallet/utils"
)

type IWalletService interface {
	Create(ctx context.Context, userID int64) (wallet *dto.Wallet, err error)
	AddFund(ctx context.Context, ID int64, fund uint64) (wallet *dto.Wallet, err error)
	SubtractFund(ctx context.Context, ID int64, fund uint64) (wallet *dto.Wallet, err error)
	GetByID(ctx context.Context, ID int64) (wallet *dto.Wallet, err error)
	GetByUserID(ctx context.Context, userID int64) (wallet *dto.Wallet, err error)
}

type WalletService struct {
	repository repository.IWalletRepository
}

func NewWalletService(repository repository.IWalletRepository) *WalletService {
	return &WalletService{repository: repository}
}

func (w *WalletService) Create(ctx context.Context, userID int64) (wallet *dto.Wallet, err error) {
	// prevent user to create two wallets
	walletEntity, err := w.repository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if walletEntity != nil {
		return nil, constants.ErrUserAlreadyHasWallet
	}

	walletEntity, err = w.repository.Create(ctx, userID)
	if err != nil {
		return nil, err
	}

	wallet = utils.WalletDTOFromEntity(walletEntity)

	return wallet, nil
}

func (w *WalletService) AddFund(ctx context.Context, ID int64, fund uint64) (wallet *dto.Wallet, err error) {
	if err := w.repository.AddFund(ctx, ID, fund); err != nil {
		return nil, err
	}

	walletEntity, err := w.repository.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	wallet = utils.WalletDTOFromEntity(walletEntity)

	return wallet, nil
}

func (w *WalletService) SubtractFund(ctx context.Context, ID int64, fund uint64) (wallet *dto.Wallet, err error) {
	walletEntity, err := w.repository.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	if walletEntity.Balance < fund {
		return nil, constants.ErrBalanceNotSufficient
	}

	if err := w.repository.SubtractFund(ctx, ID, fund); err != nil {
		return nil, err
	}

	walletEntity.Balance -= fund
	wallet = utils.WalletDTOFromEntity(walletEntity)

	return wallet, nil
}

func (w *WalletService) GetByID(ctx context.Context, ID int64) (wallet *dto.Wallet, err error) {
	walletEntity, err := w.repository.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	wallet = utils.WalletDTOFromEntity(walletEntity)

	return wallet, nil
}

func (w *WalletService) GetByUserID(ctx context.Context, userID int64) (wallet *dto.Wallet, err error) {
	walletEntity, err := w.repository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	wallet = utils.WalletDTOFromEntity(walletEntity)

	return wallet, nil
}
