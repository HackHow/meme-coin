package service

import (
	"github.com/HackHow/meme-coin/internal/model"
	"github.com/HackHow/meme-coin/internal/repository"
)

func CreateMemeCoin(coin *model.MemeCoin) error {
	return repository.CreateMemeCoin(coin)
}
