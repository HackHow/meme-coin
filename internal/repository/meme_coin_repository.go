package repository

import (
	"github.com/HackHow/meme-coin/internal/model"
	"github.com/HackHow/meme-coin/pkg/database"
)

func CreateMemeCoin(coin *model.MemeCoin) error {
	return database.DB.Create(coin).Error
}

func GetMemeCoin(id uint) (*model.MemeCoin, error) {
	var coin model.MemeCoin
	if err := database.DB.First(&coin, id).Error; err != nil {
		return nil, err
	}

	return &coin, nil
}
