package service

import (
	// "fmt"
	// "strconv"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/AsrofunNiam/lets-code-micro-service-redis/helper"
	"github.com/AsrofunNiam/lets-code-micro-service-redis/model/domain"
	"github.com/AsrofunNiam/lets-code-micro-service-redis/repository"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type RedisServiceImpl struct {
	RedisClient      *redis.Client
	Validate         *validator.Validate
	ConfigRepository repository.ConfigRepository
	DB               *gorm.DB
}

func NewRedisService(
	redisClient *redis.Client,
	validate *validator.Validate,
	configRepository repository.ConfigRepository,
	db *gorm.DB,
) RedisService {
	return &RedisServiceImpl{
		RedisClient:      redisClient,
		Validate:         validate,
		ConfigRepository: configRepository,
		DB:               db,
	}
}

func (service *RedisServiceImpl) ProcessCreateQueue(method, url, payload, key string) {
	ctx := context.Background()

	redisClient := service.RedisClient
	job := domain.JobQueue{
		Method:  method,
		URL:     url,
		Payload: payload,
	}

	jobData, err := json.Marshal(job)
	if err != nil {
		log.Fatalf("Error marshaling job data: %v", err)
	}

	// add job to list
	err = redisClient.LPush(ctx, key, jobData).Err()
	if err != nil {
		log.Fatalf("Error pushing job to queue: %v", err)
	}
	fmt.Println("Job added to queue!")

	// test redis cache
	configName, err := service.TakeRedisCacheConfig(redisClient, 1)

	helper.PanicIfError(err)
	fmt.Println(configName)

}

func (service *RedisServiceImpl) TakeRedisCacheConfig(redisClient *redis.Client, idConfig uint) (string, error) {
	ctx := context.Background()

	// Create  key by level (id config)
	key := fmt.Sprintf("config-school:%d", idConfig)

	// Check data in redis
	cachedData, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		tx := service.DB.Begin()
		defer helper.CommitOrRollback(tx)

		dataConfig := service.ConfigRepository.GetConfig(tx, &idConfig)

		// Save redis cache
		err = redisClient.Set(ctx, key, dataConfig.Name, 10*time.Minute).Err()
		if err != nil {
			return "", err
		}

		//  return value
		return dataConfig.Name, nil
	} else if err != nil {
		log.Printf("Error Process Cache Redis: %v", err)
		return "", err
	}

	// return value redis
	return cachedData, nil
}
