package handler

import (
	"github.com/gin-gonic/gin"
)

type greetingHandler struct {
}

func NewGreetingHandler() *greetingHandler {
	return &greetingHandler{}
}

func (handler *greetingHandler) HandleGreetings(context *gin.Context) {
	context.JSON(200, map[string]string{"hello": "client"})
}
