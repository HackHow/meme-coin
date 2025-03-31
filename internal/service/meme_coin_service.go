package service

import (
	"github.com/HackHow/meme-coin/internal/common"
	"github.com/HackHow/meme-coin/internal/dtos"
	"github.com/HackHow/meme-coin/internal/model"
	"github.com/HackHow/meme-coin/internal/repository"
)

func CreateMemeCoin(input dtos.CreateMemeCoinRequest) error {
	coin := &model.MemeCoin{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := repository.CreateMemeCoin(coin); err != nil {
		return &common.AppError{
			Code:    500,
			Message: "Database error",
			Data:    nil,
		}
	}

	return nil
}

func GetMemeCoin(id uint) (*model.MemeCoin, error) {
	return repository.GetMemeCoin(id)
}

func UpdateMemeCoin(id uint, input dtos.UpdateMemeCoinRequest) error {
	result := repository.UpdateMemeCoin(id, input.Description)

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
	result := repository.PokeMemeCoin(id)

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
