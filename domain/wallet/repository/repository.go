package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
	"wallet-manager/domain/constants"
	"wallet-manager/domain/wallet/entity"
)

type IWalletRepository interface {
	Create(ctx context.Context, userID int64) (wallet *entity.Wallet, err error)
	GetByID(ctx context.Context, ID int64) (wallet *entity.Wallet, err error)
	GetByUserID(ctx context.Context, userID int64) (wallet *entity.Wallet, err error)
	AddFund(ctx context.Context, ID int64, fund uint64) (err error)
	SubtractFund(ctx context.Context, ID int64, fund uint64) (err error)
}

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (w *WalletRepository) Create(ctx context.Context, userID int64) (wallet *entity.Wallet, err error) {
	result, err := w.db.ExecContext(ctx, insertWalletQuery, userID)
	if err != nil {
		log.Printf("failed to create wallet: %s\n", err)
		return nil, errors.New("failed to create wallet")
	}

	ID, err := result.LastInsertId()
	if err != nil {
		log.Printf("failed to get inserted ID: %s", err)
		return nil, errors.New("failed to create wallet")
	}

	wallet = &entity.Wallet{
		ID:        ID,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	return wallet, nil
}

func (w *WalletRepository) GetByID(ctx context.Context, ID int64) (wallet *entity.Wallet, err error) {
	result, err := w.db.QueryContext(ctx, getWalletByIDQuery, ID)
	if err != nil {
		log.Printf("failed to get wallet by ID: %s\n", err)
		return nil, errors.New("failed to get wallet by ID")
	}

	if !result.Next() {
		return nil, constants.ErrNotFound
	}

	if err := result.Scan(&wallet.ID, &wallet.Balance, &wallet.UserID, &wallet.CreatedAt, &wallet.UpdateAt); err != nil {
		log.Printf("failed to scan result: %s\n", err)
		return nil, errors.New("failed to read record")
	}

	return wallet, nil
}

func (w *WalletRepository) GetByUserID(ctx context.Context, userID int64) (wallet *entity.Wallet, err error) {
	result, err := w.db.QueryContext(ctx, getWalletByUserIDQuery, userID)
	if err != nil {
		log.Printf("failed to get wallet by user ID: %s\n", err)
		return nil, errors.New("failed to get wallet by user ID")
	}

	if !result.Next() {
		return nil, constants.ErrNotFound
	}

	if err := result.Scan(&wallet.ID, &wallet.Balance, &wallet.UserID, &wallet.CreatedAt, &wallet.UpdateAt); err != nil {
		log.Printf("failed to scan result: %s\n", err)
		return nil, errors.New("failed to read record")
	}

	return wallet, nil
}

func (w *WalletRepository) AddFund(ctx context.Context, ID int64, fund uint64) (err error) {
	_, err = w.db.ExecContext(ctx, addBalanceQuery, fund, ID)
	if err != nil {
		log.Printf("failed to add fund to wallet: %s\n", err)
		return errors.New("failed to add fund to wallet")
	}

	return nil
}

func (w *WalletRepository) SubtractFund(ctx context.Context, ID int64, fund uint64) (err error) {
	_, err = w.db.ExecContext(ctx, decreaseBalanceQuery, fund, ID)
	if err != nil {
		log.Printf("failed to decrease fund from wallet: %s\n", err)
		return errors.New("failed to decrease fund from wallet")
	}

	return nil
}
