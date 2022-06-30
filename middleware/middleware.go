package middleware

import "github.com/gin-gonic/gin"

func LoadCommonMiddleware(router gin.IRouter) {

	router.Use(CORSMiddleware())

}
