package repository

import (
	"github.com/AsrofunNiam/lets-code-micro-service_redis/helper"
	"github.com/AsrofunNiam/lets-code-micro-service_redis/model/domain"
	"gorm.io/gorm"
)

type ConfigRepositoryImpl struct {
}

func NewUserRepository() ConfigRepository {
	return &ConfigRepositoryImpl{}
}

func (repository *ConfigRepositoryImpl) GetConfig(db *gorm.DB, idConfig *uint) domain.Config {
	config := domain.Config{}

	err := db.First(&config, idConfig).Error
	helper.PanicIfError(err)

	return config
}
