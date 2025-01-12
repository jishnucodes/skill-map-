package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID        uint             `gorm:"primarykey" json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
}
