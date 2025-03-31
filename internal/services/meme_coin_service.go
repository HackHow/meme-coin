package services

import (
	"strings"

	"github.com/HackHow/meme-coin/internal/common"
	"github.com/HackHow/meme-coin/internal/dtos"
	"github.com/HackHow/meme-coin/internal/models"
	"github.com/HackHow/meme-coin/internal/repositories"
)

func CreateMemeCoin(input dtos.CreateMemeCoinRequest) error {
	coin := &models.MemeCoin{
		Name:        input.Name,
		Description: "",
	}

	if input.Description != nil {
		coin.Description = *input.Description
	}

	err := repositories.CreateMemeCoin(coin)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "UNIQUE") {
			return &common.AppError{
				Code:    409,
				Message: "Meme coin name already exists",
				Data:    nil,
			}
		}

		return &common.AppError{
			Code:    500,
			Message: "Database error",
			Data:    nil,
		}
	}

	return nil
}

func GetMemeCoin(id uint) (*models.MemeCoin, error) {
	coin, result := repositories.GetMemeCoin(id)

	if result.RowsAffected == 0 {
		return nil, &common.AppError{
			Code:    404,
			Message: "Meme coin not found",
			Data:    nil,
		}
	}

	if result.Error != nil {
		return nil, &common.AppError{
			Code:    500,
			Message: "Database error",
			Data:    nil,
		}
	}

	return coin, nil
}

func UpdateMemeCoin(id uint, input dtos.UpdateMemeCoinRequest) error {
	result := repositories.UpdateMemeCoin(id, input.Description)

	if result.RowsAffected == 0 {
		return &common.AppError{
			Code:    404,
			Message: "Meme coin not found",
			Data:    nil,
		}
	}

	if result.Error != nil {
		return &common.AppError{
			Code:    500,
			Message: "Database error",
			Data:    nil,
		}
	}

	return nil
}

func DeleteMemeCoin(id uint) error {
	result := repositories.DeleteMemeCoin(id)

	if result.RowsAffected == 0 {
		return &common.AppError{
			Code:    404,
			Message: "Meme coin not found",
			Data:    nil,
		}
	}

	if result.Error != nil {
		return &common.AppError{
			Code:    500,
			Message: "Database error",
			Data:    nil,
		}
	}

	return nil
}

func PokeMemeCoin(id uint) error {
	result := repositories.PokeMemeCoin(id)

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
