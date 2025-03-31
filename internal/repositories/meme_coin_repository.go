package repositories

import (
	"github.com/HackHow/meme-coin/internal/models"
	"github.com/HackHow/meme-coin/pkg/database"
	"gorm.io/gorm"
)

func CreateMemeCoin(coin *models.MemeCoin) error {
	return database.DB.Create(coin).Error
}

func GetMemeCoin(id uint) (*models.MemeCoin, *gorm.DB) {
	var coin models.MemeCoin
	result := database.DB.First(&coin, id)

	return &coin, result
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
