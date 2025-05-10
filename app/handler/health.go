package handler

import (
	"github.com/go-playground/validator/v10"
	"go-gin-pattern/app/service"
)

type HealthHandler struct {
	Service   *service.HealthService
	Validator *validator.Validate
}

func NewHealthHandler(service *service.HealthService, validator *validator.Validate) *HealthHandler {
	return &HealthHandler{
		service,
		validator,
	}
}
