package core

import (
	"github.com/ashkanamani/url-shortener/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	Config *config.Config
	Router *gin.Engine
	Logger *logrus.Logger
	DB     *gorm.DB
}
