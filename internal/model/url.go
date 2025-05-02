package model

import (
	"gorm.io/gorm"
	"time"
)

type Url struct {
	ID         string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID     string         `gorm:"type:uuid;not null"`
	Original   string         `gorm:"type:text;not null"`
	ShortCode  string         `gorm:"type:varchar(10);unique;not null"`
	VisitCount int            `gorm:"default:0"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
