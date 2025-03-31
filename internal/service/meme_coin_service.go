package service

import (
	"github.com/HackHow/meme-coin/internal/common"
	"github.com/HackHow/meme-coin/internal/model"
	"github.com/HackHow/meme-coin/internal/repository"
)

func CreateMemeCoin(coin *model.MemeCoin) error {
	return repository.CreateMemeCoin(coin)
}

func GetMemeCoin(id uint) (*model.MemeCoin, error) {
	return repository.GetMemeCoin(id)
}

func UpdateMemeCoin(id uint, description string) error {
	result := repository.UpdateMemeCoin(id, description)

	if result.Error != nil {
		return &common.AppError{
			Code:    500,
			Message: "Database error",
			Data:    nil,
		}
	}

	if result.RowsAffected == 0 {
		return &common.AppError{
			Code:    404,
			Message: "Meme coin not found",
			Data:    nil,
		}
	}

	return nil
}

func DeleteMemeCoin(id uint) error {
	result := repository.DeleteMemeCoin(id)

	if result.Error != nil {
		return &common.AppError{
			Code:    500,
			Message: "Database error",
			Data:    nil,
		}
	}

	if result.RowsAffected == 0 {
		return &common.AppError{
			Code:    404,
			Message: "Meme coin not found",
			Data:    nil,
		}
	}

	return nil
}

func PokeMemeCoin(id uint) error {
	return repository.PokeMemeCoin(id)
}
