package service

import (
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
	return repository.UpdateMemeCoin(id, description)
}

func DeleteMemeCoin(id uint) error {
	return repository.DeleteMemeCoin(id)
}
