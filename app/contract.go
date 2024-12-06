package app

import (
	"github.com/QBC8-Go-Group2/questionnaire/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type App interface {
	Config() config.Config
	DB() *gorm.DB
	Redis() *redis.Client
}
