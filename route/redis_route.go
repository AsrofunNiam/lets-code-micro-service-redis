package route

import (
	"github.com/AsrofunNiam/lets-code-micro-service_redis/controller"
	"github.com/AsrofunNiam/lets-code-micro-service_redis/repository"
	"github.com/AsrofunNiam/lets-code-micro-service_redis/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RedisRoute(router *gin.Engine, db *gorm.DB, redisClient *redis.Client, validate *validator.Validate) {

	creditNoteService := service.NewRedisService(
		redisClient,
		validate,
		repository.NewUserRepository(),
		db,
	)
	creditNoteController := controller.NewRedisController(creditNoteService)

	router.POST("/redis/process/create/queue", creditNoteController.ProcessCreateQueue)

}
