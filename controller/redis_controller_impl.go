package controller

import (
	"net/http"

	"github.com/AsrofunNiam/lets-code-micro-service_redis/helper"
	"github.com/AsrofunNiam/lets-code-micro-service_redis/model/web"
	"github.com/AsrofunNiam/lets-code-micro-service_redis/service"
	"github.com/gin-gonic/gin"
)

type RedisControllerImpl struct {
	RedisService service.RedisService
}

func NewRedisController(redisService service.RedisService) RedisController {
	return &RedisControllerImpl{
		RedisService: redisService,
	}
}

func (controller *RedisControllerImpl) ProcessCreateQueue(c *gin.Context) {
	request := web.JobsQueueCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	controller.RedisService.ProcessCreateQueue(request.Method, request.URL, request.Payload, request.Key)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Process create redis queue successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}
