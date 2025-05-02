package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email        string         `gorm:"type:text;unique;not null"`
	PasswordHash string         `gorm:"type:text;not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
