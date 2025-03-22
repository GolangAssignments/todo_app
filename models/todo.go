package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         `gorm:"size:255"`
	Completed bool           `gorm:"default:false"`
}
