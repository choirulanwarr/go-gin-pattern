package config

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-gin-pattern/app/handler"
	"go-gin-pattern/app/router"
	"go-gin-pattern/app/service"
	"gorm.io/gorm"
)

type AppConfig struct {
	Config    *viper.Viper
	DB        *gorm.DB
	Server    *gin.Engine
	Validator *validator.Validate
	Logger    *logrus.Logger
}

func InitConfig(config *AppConfig) {
	// Health
	healthService := service.NewHealthService(config.Config)
	healthHandler := handler.NewHealthHandler(healthService, config.Validator)

	// Routers
	routeConfig := router.Config{
		Server:        config.Server,
		Config:        config.Config,
		DB:            config.DB,
		Logger:        config.Logger,
		HealthHandler: healthHandler,
	}

	routeConfig.Init()
}
