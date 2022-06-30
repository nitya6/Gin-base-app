package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	HandleGreetings(context *gin.Context)
	HandleSomething(context *gin.Context)
}
