package service

import "github.com/spf13/viper"

type HealthService struct {
	Viper *viper.Viper
}

func NewHealthService(viper *viper.Viper) *HealthService {
	return &HealthService{
		viper,
	}
}
