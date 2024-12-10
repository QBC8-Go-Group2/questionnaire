package app

import (
	"github.com/QBC8-Go-Group2/questionnaire/config"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
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
	a := app{config: config}

	if err := a.setDB(); err != nil {
		return nil, err
	}
	if err := a.setRedis(); err != nil {
		return nil, err
	}

	if err := a.migrationDB(); err != nil {
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
	var dbCfg = a.config.DB
	var db, err = mysql.NewMySqlGormConnection(mysql.DBConnectionConfig{
		Host:   dbCfg.Host,
		Port:   dbCfg.Port,
		User:   dbCfg.User,
		Pass:   dbCfg.Password,
		Dbname: dbCfg.Database,
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

func (a *app) migrationDB() error {
	if err := a.db.AutoMigrate(&types.User{}); err != nil {
		return err
	}
	if err := a.db.AutoMigrate(&types.Questionnaire{}); err != nil {
		return err
	}
	if err := a.db.AutoMigrate(&types.Question{}); err != nil {
		return err
	}
	if err := a.db.AutoMigrate(&types.Option{}); err != nil {
		return err
	}
	if err := a.db.AutoMigrate(&types.Response{}); err != nil {
		return err
	}
	if err := a.db.AutoMigrate(&types.Media{}); err != nil {
		return err
	}
	return nil
}
