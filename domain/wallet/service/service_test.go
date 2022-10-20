package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"wallet-manager/domain/constants"
	"wallet-manager/domain/wallet/dto"
	"wallet-manager/domain/wallet/entity"
	mocks "wallet-manager/mocks/domain/wallet/repository"
)

type ServiceTestSuite struct {
	suite.Suite
}

func (s *ServiceTestSuite) TestCreate() {
	testCases := []struct {
		ctx                  context.Context
		userID               int64
		expectedWalletDTO    *dto.Wallet
		expectedWalletEntity *entity.Wallet
		expectedError        error
	}{
		{
			ctx:    context.Background(),
			userID: 1,
			expectedWalletDTO: &dto.Wallet{
				ID:      1,
				Balance: 0,
				UserID:  1,
			},
			expectedWalletEntity: &entity.Wallet{
				ID:      1,
				Balance: 0,
				UserID:  1,
			},
			expectedError: nil,
		},
		{
			ctx:    context.Background(),
			userID: 1,
			expectedWalletEntity: &entity.Wallet{
				ID:      1,
				Balance: 0,
				UserID:  1,
			},
			expectedError: constants.ErrUserAlreadyHasWallet,
		},
		{
			ctx:           context.Background(),
			userID:        1,
			expectedError: errors.New("failed to get wallet by user ID"),
		},
	}

	for _, tc := range testCases {
		if tc.expectedError != nil {
			if errors.Is(tc.expectedError, constants.ErrUserAlreadyHasWallet) {
				repositoryMock := mocks.IWalletRepository{}
				repositoryMock.On("Create", mock.Anything, tc.userID).Return(nil, tc.expectedError)
				repositoryMock.On("GetByUserID", mock.Anything, tc.userID).Return(tc.expectedWalletEntity, nil)
				walletService := NewWalletService(&repositoryMock)
				_, err := walletService.Create(tc.ctx, tc.userID)
				assert.Equal(s.T(), tc.expectedError, err)
			} else {
				repositoryMock := mocks.IWalletRepository{}
				repositoryMock.On("Create", mock.Anything, tc.userID).Return(nil, tc.expectedError)
				repositoryMock.On("GetByUserID", mock.Anything, tc.userID).Return(nil, tc.expectedError)
				walletService := NewWalletService(&repositoryMock)
				_, err := walletService.Create(tc.ctx, tc.userID)
				assert.Equal(s.T(), tc.expectedError, err)
			}

		} else {
			repositoryMock := mocks.IWalletRepository{}
			repositoryMock.On("Create", mock.Anything, tc.userID).Return(tc.expectedWalletEntity, nil)
			repositoryMock.On("GetByUserID", mock.Anything, tc.userID).Return(nil, tc.expectedError)
			walletService := NewWalletService(&repositoryMock)
			walletDTO, err := walletService.Create(tc.ctx, tc.userID)
			assert.Equal(s.T(), tc.expectedError, err)
			assert.Equal(s.T(), walletDTO, tc.expectedWalletDTO)
		}
	}
}

func (s *ServiceTestSuite) TestAddFund() {
	testCases := []struct {
		ctx                  context.Context
		id                   int64
		fund                 uint64
		expectedWalletEntity *entity.Wallet
		expectedWalletDTO    *dto.Wallet
		expectedError        error
	}{
		{
			ctx:  context.Background(),
			id:   1,
			fund: 100,
			expectedWalletEntity: &entity.Wallet{
				ID:      1,
				Balance: 100,
			},
			expectedWalletDTO: &dto.Wallet{
				ID:      1,
				Balance: 100,
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		repositoryMock := mocks.IWalletRepository{}
		repositoryMock.On("AddFund", tc.ctx, tc.id, tc.fund).Return(tc.expectedError)
		repositoryMock.On("GetByID", tc.ctx, tc.id).Return(tc.expectedWalletEntity, nil)
		walletService := NewWalletService(&repositoryMock)
		walletDTO, err := walletService.AddFund(tc.ctx, tc.id, tc.fund)
		assert.Equal(s.T(), tc.expectedError, err)
		assert.Equal(s.T(), tc.expectedWalletDTO.ID, walletDTO.ID)
		assert.Equal(s.T(), tc.expectedWalletDTO.Balance, walletDTO.Balance)
	}
}

func (s *ServiceTestSuite) TestSubtractFund() {
	testCases := []struct {
		ctx                  context.Context
		id                   int64
		fund                 uint64
		expectedWalletEntity *entity.Wallet
		expectedWalletDTO    *dto.Wallet
		expectedError        error
	}{
		{
			ctx:  context.Background(),
			id:   1,
			fund: 100,
			expectedWalletEntity: &entity.Wallet{
				ID:      1,
				Balance: 100,
			},
			expectedWalletDTO: &dto.Wallet{
				ID:      1,
				Balance: 0,
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		repositoryMock := mocks.IWalletRepository{}
		repositoryMock.On("SubtractFund", tc.ctx, tc.id, tc.fund).Return(tc.expectedError)
		repositoryMock.On("GetByID", tc.ctx, tc.id).Return(tc.expectedWalletEntity, nil)
		walletService := NewWalletService(&repositoryMock)
		walletDTO, err := walletService.SubtractFund(tc.ctx, tc.id, tc.fund)
		assert.Equal(s.T(), tc.expectedError, err)
		assert.Equal(s.T(), tc.expectedWalletDTO.ID, walletDTO.ID)
		assert.Equal(s.T(), tc.expectedWalletDTO.Balance, walletDTO.Balance)
	}
}

func (s *ServiceTestSuite) TestGetByID() {
	testCases := []struct {
		ctx                  context.Context
		id                   int64
		expectedWalletEntity *entity.Wallet
		expectedWalletDTO    *dto.Wallet
		expectedError        error
	}{
		{
			ctx: context.Background(),
			id:  1,
			expectedWalletEntity: &entity.Wallet{
				ID:      1,
				Balance: 0,
				UserID:  1,
			},
			expectedWalletDTO: &dto.Wallet{
				ID:      1,
				Balance: 0,
				UserID:  1,
			},
			expectedError: nil,
		},
		{
			ctx:                  context.Background(),
			id:                   1,
			expectedWalletEntity: nil,
			expectedError:        constants.ErrNotFound,
		},
		{
			ctx:                  context.Background(),
			id:                   1,
			expectedWalletEntity: nil,
			expectedError:        errors.New("failed to get wallet by ID"),
		},
		{
			ctx:                  context.Background(),
			id:                   1,
			expectedWalletEntity: nil,
			expectedError:        errors.New("failed to read record"),
		},
	}

	for _, tc := range testCases {
		repositoryMock := mocks.IWalletRepository{}
		repositoryMock.On("GetByID", tc.ctx, tc.id).Return(tc.expectedWalletEntity, tc.expectedError)

		walletService := NewWalletService(&repositoryMock)
		walletDTO, err := walletService.GetByID(tc.ctx, tc.id)
		assert.Equal(s.T(), tc.expectedError, err)
		if tc.expectedError != nil {
			assert.Equal(s.T(), tc.expectedError, err)
		} else {
			assert.Equal(s.T(), tc.expectedWalletDTO.ID, walletDTO.ID)
			assert.Equal(s.T(), tc.expectedWalletDTO.Balance, walletDTO.Balance)
			assert.Equal(s.T(), tc.expectedWalletDTO.UserID, walletDTO.UserID)
		}
	}
}

func (s *ServiceTestSuite) TestGetByUserID() {
	testCases := []struct {
		ctx                  context.Context
		id                   int64
		expectedWalletEntity *entity.Wallet
		expectedWalletDTO    *dto.Wallet
		expectedError        error
	}{
		{
			ctx: context.Background(),
			id:  1,
			expectedWalletEntity: &entity.Wallet{
				ID:      1,
				Balance: 0,
				UserID:  1,
			},
			expectedWalletDTO: &dto.Wallet{
				ID:      1,
				Balance: 0,
				UserID:  1,
			},
			expectedError: nil,
		},
		{
			ctx:                  context.Background(),
			id:                   1,
			expectedWalletEntity: nil,
			expectedError:        constants.ErrNotFound,
		},
		{
			ctx:                  context.Background(),
			id:                   1,
			expectedWalletEntity: nil,
			expectedError:        errors.New("failed to get wallet by user ID"),
		},
		{
			ctx:                  context.Background(),
			id:                   1,
			expectedWalletEntity: nil,
			expectedError:        errors.New("failed to read record"),
		},
	}

	for _, tc := range testCases {
		repositoryMock := mocks.IWalletRepository{}
		repositoryMock.On("GetByUserID", tc.ctx, tc.id).Return(tc.expectedWalletEntity, tc.expectedError)

		walletService := NewWalletService(&repositoryMock)
		walletDTO, err := walletService.GetByUserID(tc.ctx, tc.id)
		assert.Equal(s.T(), tc.expectedError, err)
		if tc.expectedError != nil {
			assert.Equal(s.T(), tc.expectedError, err)
		} else {
			assert.Equal(s.T(), tc.expectedWalletDTO.ID, walletDTO.ID)
			assert.Equal(s.T(), tc.expectedWalletDTO.Balance, walletDTO.Balance)
			assert.Equal(s.T(), tc.expectedWalletDTO.UserID, walletDTO.UserID)
		}
	}
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
