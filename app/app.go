package app

import (
	"github.com/QBC8-Go-Group2/questionnaire/config"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/mysql"
	redispkg "github.com/QBC8-Go-Group2/questionnaire/pkg/redis"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type app struct {
	config config.Config
	db     *gorm.DB
	redis  *redis.Client
}

func (a app) Config() config.Config {
	return a.config
}

func (a app) DB() *gorm.DB {
	return a.db
}

func (a app) Redis() *redis.Client {
	return a.redis
}

func NewApp(config config.Config) (App, error) {
	a := &app{config: config}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	if err := a.setRedis(); err != nil {
		return nil, err
	}

	return a, nil
}

func MustNewApp(config config.Config) App {
	a, err := NewApp(config)
	if err != nil {
		panic(err)
	}
	return a
}

func (a *app) setDB() error {
	db, err := mysql.NewMySqlGormConnection(mysql.DBConnectionConfig{
		Host:   a.config.DB.Host,
		Port:   a.config.DB.Port,
		User:   a.config.DB.User,
		Pass:   a.config.DB.Password,
		Dbname: a.config.DB.Database,
	})
	if err != nil {
		return err
	}
	a.db = db
	return nil
}

func (a *app) setRedis() error {
	client, err := redispkg.NewRedisClient(redispkg.Config{
		Host: a.config.Redis.Host,
		Port: a.config.Redis.Port,
	})
	if err != nil {
		return err
	}
	a.redis = client
	return nil
}
