package middleware

import (
	"base-app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Authentication middleware to validate request
func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		Logger := utils.Logger

		Logger.Info("Intercepting auth route", zap.String("URL", ctx.Request.URL.Path))

		ctx.Next()
	}

}
