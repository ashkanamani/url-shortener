package postgres

import (
	"github.com/ashkanamani/url-shortener/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(addr string, logger *logrus.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(addr), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&model.Url{}, &model.User{}); err != nil {
		logger.Fatalf("failed to migrate database: %v", err)
	}
	logger.Infof("successfully connected to database")
	return db
}
