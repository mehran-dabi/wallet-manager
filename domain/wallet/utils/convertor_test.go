package utils

import (
	"testing"
	"time"
	"wallet-manager/domain/wallet/dto"
	"wallet-manager/domain/wallet/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConvertorTestSuite struct {
	suite.Suite
}

func (u *ConvertorTestSuite) TestWalletDTOFromEntity() {
	currentTime := time.Now()
	testCases := []struct {
		walletEntity      *entity.Wallet
		expectedWalletDTO *dto.Wallet
	}{
		{
			walletEntity: &entity.Wallet{
				ID:        1,
				Balance:   0,
				UserID:    1,
				CreatedAt: currentTime,
				UpdatedAt: currentTime,
			},
			expectedWalletDTO: &dto.Wallet{
				ID:        1,
				Balance:   0,
				UserID:    1,
				CreatedAt: currentTime,
				UpdatedAt: currentTime,
			},
		},
		{
			walletEntity: nil,
		},
	}

	for _, tc := range testCases {
		if tc.walletEntity != nil {
			walletEntity := WalletDTOFromEntity(tc.walletEntity)
			assert.Equal(u.T(), tc.expectedWalletDTO, walletEntity)
		} else {
			assert.Panics(u.T(), func() { WalletDTOFromEntity(tc.walletEntity) })
		}
	}
}

func (u *ConvertorTestSuite) WalletEntityFromDTO() {
	currentTime := time.Now()
	testCases := []struct {
		walletDTO            *dto.Wallet
		expectedWalletEntity *entity.Wallet
	}{
		{
			walletDTO: &dto.Wallet{
				ID:        1,
				Balance:   0,
				UserID:    1,
				CreatedAt: currentTime,
				UpdatedAt: currentTime,
			},
			expectedWalletEntity: &entity.Wallet{
				ID:        1,
				Balance:   0,
				UserID:    1,
				CreatedAt: currentTime,
				UpdatedAt: currentTime,
			},
		},
		{
			walletDTO: nil,
		},
	}

	for _, tc := range testCases {
		if tc.walletDTO != nil {
			walletEntity := WalletEntityFromDTO(tc.walletDTO)
			assert.Equal(u.T(), tc.expectedWalletEntity, walletEntity)
		} else {
			assert.Panics(u.T(), func() { WalletEntityFromDTO(tc.walletDTO) })
		}
	}
}

func TestConvertorTestSuite(t *testing.T) {
	suite.Run(t, new(ConvertorTestSuite))
}
