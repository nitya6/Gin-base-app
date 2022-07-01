package middleware

import "github.com/gin-gonic/gin"

// load common middleware for all routes
func LoadCommonMiddleware(router gin.IRouter) {

	router.Use(CORSMiddleware())

}
