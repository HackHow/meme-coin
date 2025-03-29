package repository

import (
	"github.com/HackHow/meme-coin/internal/model"
	"github.com/HackHow/meme-coin/pkg/database"
)

func CreateMemeCoin(coin *model.MemeCoin) error {
	return database.DB.Create(coin).Error
}
