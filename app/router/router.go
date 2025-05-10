package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-gin-pattern/app/handler"
	"gorm.io/gorm"
)

type Config struct {
	Server        *gin.Engine
	Config        *viper.Viper
	DB            *gorm.DB
	Logger        *logrus.Logger
	HealthHandler *handler.HealthHandler
}

func (c *Config) Init() {
	initGlobalRoutes(c)
}
