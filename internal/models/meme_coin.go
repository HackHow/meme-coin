package models

import (
	"time"
)

type MemeCoin struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`   // The unique identifier of a meme coin.
	Name            string    `gorm:"unique;not null" json:"name"`          // The name of the meme coin. Must be unique across all coins and cannot be empty.
	Description     string    `gorm:"type:varchar(255)" json:"description"` // A brief description of the meme coin.
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`     // The timestamp when the meme coin was created.
	PopularityScore int       `gorm:"default:0" json:"popularity_score"`    // A score representing the popularity of the meme coin.
}
