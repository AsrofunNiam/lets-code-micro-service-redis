package repository

import (
	"github.com/AsrofunNiam/lets-code-micro-service-redis/model/domain"
	"gorm.io/gorm"
)

type ConfigRepository interface {
	GetConfig(db *gorm.DB, idConfig *uint) domain.Config
}
