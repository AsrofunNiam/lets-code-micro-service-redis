package controller

import (
	"github.com/gin-gonic/gin"
)

type RedisController interface {
	ProcessCreateQueue(context *gin.Context)
}
