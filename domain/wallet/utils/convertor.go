package utils

import (
	"wallet-manager/domain/wallet/dto"
	"wallet-manager/domain/wallet/entity"
)

func WalletDTOFromEntity(entity *entity.Wallet) *dto.Wallet {
	return &dto.Wallet{
		ID:        entity.ID,
		Balance:   entity.Balance,
		UserID:    entity.UserID,
		CreatedAt: entity.CreatedAt,
		UpdateAt:  entity.UpdateAt,
	}
}

func WalletEntityFromDTO(dto *dto.Wallet) *entity.Wallet {
	return &entity.Wallet{
		ID:        dto.ID,
		Balance:   dto.Balance,
		UserID:    dto.UserID,
		CreatedAt: dto.CreatedAt,
		UpdateAt:  dto.UpdateAt,
	}
}
