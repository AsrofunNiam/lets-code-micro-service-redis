package app

import (
	route "github.com/AsrofunNiam/lets-code-micro-service_redis/route"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewRouter(redisClient *redis.Client, db *gorm.DB, validate *validator.Validate) *gin.Engine {

	router := gin.New()
	router.UseRawPath = true

	route.RedisRoute(router, db, redisClient, validate)

	return router
}
