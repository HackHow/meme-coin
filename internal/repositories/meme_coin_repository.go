package repositories

import (
	"github.com/HackHow/meme-coin/internal/models"
	"github.com/HackHow/meme-coin/pkg/database"
	"gorm.io/gorm"
)

func CreateMemeCoin(coin *models.MemeCoin) error {
	return database.DB.Create(coin).Error
}

func GetMemeCoin(id uint) (*models.MemeCoin, error) {
	var coin models.MemeCoin
	if err := database.DB.First(&coin, id).Error; err != nil {
		return nil, err
	}

	return &coin, nil
}

func UpdateMemeCoin(id uint, description string) *gorm.DB {
	return database.DB.Model(&models.MemeCoin{}).
		Where("id = ?", id).
		Update("description", description)
}

func DeleteMemeCoin(id uint) *gorm.DB {
	return database.DB.Delete(&models.MemeCoin{}, id)
}

func PokeMemeCoin(id uint) *gorm.DB {
	return database.DB.Model(&models.MemeCoin{}).
		Where("id = ?", id).
		UpdateColumn("popularity_score", gorm.Expr("popularity_score + ?", 1))
}
