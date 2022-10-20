package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
	"wallet-manager/domain/constants"
	"wallet-manager/domain/wallet/entity"
	mocks "wallet-manager/mocks/infrastructure/database"
)

type RepositoryTestSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func (r *RepositoryTestSuite) TestCreate() {
	testCases := []struct {
		userID               int64
		ctx                  context.Context
		expectedWalletEntity *entity.Wallet
		expectedError        error
	}{
		{
			userID: 1,
			ctx:    context.Background(),
			expectedWalletEntity: &entity.Wallet{
				ID:      1,
				Balance: 0,
				UserID:  1,
			},
			expectedError: nil,
		},
		{
			userID: 1,
			ctx:    context.Background(),
			expectedWalletEntity: &entity.Wallet{
				ID:      1,
				Balance: 0,
				UserID:  1,
			},
			expectedError: errors.New("failed to create wallet"),
		},
	}

	r.db, r.mock = mocks.NewDBMock()
	walletRepository := NewWalletRepository(r.db)

	for _, tc := range testCases {
		if tc.expectedError != nil {
			r.mock.ExpectExec("INSERT INTO").
				WithArgs(tc.userID).
				WillReturnError(tc.expectedError)
			_, err := walletRepository.Create(tc.ctx, tc.userID)
			assert.Equal(r.T(), tc.expectedError, err)
		} else {
			r.mock.ExpectExec("INSERT INTO").
				WithArgs(tc.userID).
				WillReturnError(tc.expectedError).
				WillReturnResult(sqlmock.NewResult(1, 1))
			wallet, err := walletRepository.Create(tc.ctx, tc.userID)
			assert.Equal(r.T(), tc.expectedError, err)
			assert.Equal(r.T(), tc.expectedWalletEntity.ID, wallet.ID)
			assert.Equal(r.T(), tc.expectedWalletEntity.Balance, wallet.Balance)
			assert.Equal(r.T(), tc.expectedWalletEntity.UserID, wallet.UserID)
		}
	}
}

func (r *RepositoryTestSuite) TestAddFund() {
	testCases := []struct {
		ctx           context.Context
		ID            int64
		fund          uint64
		expectedError error
	}{
		{
			ctx:           context.Background(),
			ID:            1,
			fund:          100,
			expectedError: nil,
		},
		{
			ctx:           context.Background(),
			ID:            1,
			fund:          100,
			expectedError: errors.New("failed to add fund to wallet"),
		},
	}

	r.db, r.mock = mocks.NewDBMock()
	walletRepository := NewWalletRepository(r.db)

	for _, tc := range testCases {
		if tc.expectedError != nil {
			r.mock.ExpectExec("UPDATE wallets").
				WithArgs(tc.fund, tc.ID).
				WillReturnError(tc.expectedError)
		} else {
			r.mock.ExpectExec("UPDATE wallets").
				WithArgs(tc.fund, tc.ID).
				WillReturnResult(sqlmock.NewResult(0, 1))
		}
		err := walletRepository.AddFund(tc.ctx, tc.ID, tc.fund)
		assert.Equal(r.T(), tc.expectedError, err)
	}
}

func (r *RepositoryTestSuite) TestSubtractFund() {
	testCases := []struct {
		ctx           context.Context
		ID            int64
		fund          uint64
		expectedError error
	}{
		{
			ctx:           context.Background(),
			ID:            1,
			fund:          100,
			expectedError: nil,
		},
		{
			ctx:           context.Background(),
			ID:            1,
			fund:          100,
			expectedError: errors.New("failed to subtract fund from wallet"),
		},
	}

	r.db, r.mock = mocks.NewDBMock()
	walletRepository := NewWalletRepository(r.db)

	for _, tc := range testCases {
		if tc.expectedError != nil {
			r.mock.ExpectExec("UPDATE wallets").
				WithArgs(tc.fund, tc.ID).
				WillReturnError(tc.expectedError)
		} else {
			r.mock.ExpectExec("UPDATE wallets").
				WithArgs(tc.fund, tc.ID).
				WillReturnResult(sqlmock.NewResult(0, 1))
		}
		err := walletRepository.SubtractFund(tc.ctx, tc.ID, tc.fund)
		assert.Equal(r.T(), tc.expectedError, err)
	}
}

func (r *RepositoryTestSuite) TestGetByID() {
	testCases := []struct {
		id                   int64
		ctx                  context.Context
		expectedWalletEntity *entity.Wallet
		expectedError        error
	}{
		{
			id:  1,
			ctx: context.Background(),
			expectedWalletEntity: &entity.Wallet{
				ID:        1,
				Balance:   0,
				UserID:    1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedError: nil,
		},
		{
			id:            1,
			ctx:           context.Background(),
			expectedError: constants.ErrNotFound,
		},
		{
			id:            1,
			ctx:           context.Background(),
			expectedError: errors.New("failed to get wallet by ID"),
		},
	}

	r.db, r.mock = mocks.NewDBMock()
	walletRepository := NewWalletRepository(r.db)

	for _, tc := range testCases {
		if tc.expectedError != nil {
			if errors.Is(tc.expectedError, constants.ErrNotFound) {
				rows := r.mock.NewRows([]string{"id", "balance", "user_id", "created_at", "updated_at"})
				r.mock.ExpectQuery("SELECT id, balance, user_id, created_at, updated_at FROM wallets").
					WithArgs(tc.id).
					WillReturnRows(rows)
				_, err := walletRepository.GetByID(tc.ctx, tc.id)
				assert.Equal(r.T(), tc.expectedError, err)
			} else {
				r.mock.ExpectQuery("SELECT id, balance, user_id, created_at, updated_at FROM wallets").
					WithArgs(tc.id).
					WillReturnError(constants.ErrNotFound)
				_, err := walletRepository.GetByID(tc.ctx, tc.id)
				assert.Equal(r.T(), tc.expectedError, err)
			}
		} else {
			rows := r.mock.NewRows([]string{"id", "balance", "user_id", "created_at", "updated_at"}).
				AddRow(
					tc.expectedWalletEntity.ID,
					tc.expectedWalletEntity.Balance,
					tc.expectedWalletEntity.UserID,
					tc.expectedWalletEntity.CreatedAt,
					tc.expectedWalletEntity.UpdatedAt,
				)

			r.mock.ExpectQuery("SELECT id, balance, user_id, created_at, updated_at FROM wallets").
				WithArgs(tc.id).
				WillReturnRows(rows)
			walletEntity, err := walletRepository.GetByID(tc.ctx, tc.id)
			assert.Equal(r.T(), tc.expectedError, err)
			assert.Equal(r.T(), tc.expectedWalletEntity, walletEntity)
		}
	}
}

func (r *RepositoryTestSuite) TestGetByUserID() {
	testCases := []struct {
		userID               int64
		ctx                  context.Context
		expectedWalletEntity *entity.Wallet
		expectedError        error
	}{
		{
			userID: 1,
			ctx:    context.Background(),
			expectedWalletEntity: &entity.Wallet{
				ID:        1,
				Balance:   0,
				UserID:    1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedError: nil,
		},
		{
			userID:        1,
			ctx:           context.Background(),
			expectedError: constants.ErrNotFound,
		},
		{
			userID:        1,
			ctx:           context.Background(),
			expectedError: errors.New("failed to get wallet by user ID"),
		},
	}

	r.db, r.mock = mocks.NewDBMock()
	walletRepository := NewWalletRepository(r.db)

	for _, tc := range testCases {
		if tc.expectedError != nil {
			if errors.Is(tc.expectedError, constants.ErrNotFound) {
				rows := r.mock.NewRows([]string{"id", "balance", "user_id", "created_at", "updated_at"})
				r.mock.ExpectQuery("SELECT id, balance, user_id, created_at, updated_at FROM wallets").
					WithArgs(tc.userID).
					WillReturnRows(rows)
				_, err := walletRepository.GetByUserID(tc.ctx, tc.userID)
				assert.Equal(r.T(), tc.expectedError, err)
			} else {
				r.mock.ExpectQuery("SELECT id, balance, user_id, created_at, updated_at FROM wallets").
					WithArgs(tc.userID).
					WillReturnError(constants.ErrNotFound)
				_, err := walletRepository.GetByUserID(tc.ctx, tc.userID)
				assert.Equal(r.T(), tc.expectedError, err)
			}
		} else {
			rows := r.mock.NewRows([]string{"id", "balance", "user_id", "created_at", "updated_at"}).
				AddRow(
					tc.expectedWalletEntity.ID,
					tc.expectedWalletEntity.Balance,
					tc.expectedWalletEntity.UserID,
					tc.expectedWalletEntity.CreatedAt,
					tc.expectedWalletEntity.UpdatedAt,
				)

			r.mock.ExpectQuery("SELECT id, balance, user_id, created_at, updated_at FROM wallets").
				WithArgs(tc.userID).
				WillReturnRows(rows)
			walletEntity, err := walletRepository.GetByUserID(tc.ctx, tc.userID)
			assert.Equal(r.T(), tc.expectedError, err)
			assert.Equal(r.T(), tc.expectedWalletEntity, walletEntity)
		}
	}
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
