package repository

import (
	"github.com/HackHow/meme-coin/internal/model"
	"github.com/HackHow/meme-coin/pkg/database"
	"gorm.io/gorm"
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

func UpdateMemeCoin(id uint, description string) *gorm.DB {
	return database.DB.Model(&model.MemeCoin{}).
		Where("id = ?", id).
		Update("description", description)
}

func DeleteMemeCoin(id uint) *gorm.DB {
	return database.DB.Delete(&model.MemeCoin{}, id)
}

func PokeMemeCoin(id uint) error {
	return database.DB.Model(&model.MemeCoin{}).
		Where("id = ?", id).
		UpdateColumn("popularity_score", gorm.Expr("popularity_score + ?", 1)).Error
}
