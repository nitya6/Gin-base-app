package api

import (
	"base-app/middleware"
	"base-app/pkg/handler"
	"github.com/gin-gonic/gin"
)

// Applications rest interface
func LoadRoutes(router gin.IRouter) {

	v1 := router.Group("/api/v1")
	v1Authorized := v1.Group("authorized")

	v1.GET("/greetings", handler.NewGreetingHandler().HandleGreetings)

	v1Authorized.GET("", middleware.AuthMiddleware(), handler.NewGreetingHandler().HandleGreetings)
}
