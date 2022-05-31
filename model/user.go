package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Users struct {
	gorm.Model
	ID        int       `gorm:"id pk autoincr"`
	Name      string    `gorm:"name"`
	Address   string    `gorm:"address"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

type UserInput struct {
	gorm.Model
	Name    string `json:"name" binding:"required" gorm:"name"`
	Address string `json:"address" binding:"required" gorm:"address"`
}
