package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AsrofunNiam/lets-code-micro-service_redis/app"
	c "github.com/AsrofunNiam/lets-code-micro-service_redis/configuration"
	"github.com/AsrofunNiam/lets-code-micro-service_redis/helper"
	"github.com/go-playground/validator/v10"
)

func main() {
	configuration, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	port := configuration.Port
	db := app.ConnectDatabase(configuration.User, configuration.Host, configuration.Password, configuration.PortDB, configuration.Db)
	redisClient := app.ConnectClientCRedis(configuration.RedisHost, configuration.RedisPort, configuration.RedisPassword)

	validate := validator.New()
	router := app.NewRouter(redisClient, db, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Server is running on port %s", port)

	//  run always take redis
	listKey := helper.ListKey()
	for _, key := range listKey.Keys {
		fmt.Println("running key: ", key)
		go helper.TakeRedisQueue(redisClient, key)
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
