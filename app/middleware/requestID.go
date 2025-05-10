package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin-pattern/app/constant"
	"go-gin-pattern/app/helper"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiID := helper.GenerateApiCallID()
		c.Set(constant.RequestIDKey, apiID)

		c.Next()
	}
}
