package app

import (
	"github.com/QBC8-Go-Group2/questionnaire/config"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/mysql"
	"gorm.io/gorm"
)

type app struct {
	config config.Config
	db     *gorm.DB
}

func (a app) Config() config.Config {
	return a.config
}

func NewApp(config config.Config) (App, error) {
	a := app{config: config}

	if err := a.setDB(); err != nil {
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

func (a app) setDB() error {
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

func (a app) migrationDB() error {
	if err := a.db.Table("users").AutoMigrate(&types.User{}); err != nil {
		return err
	}
	if err := a.db.Table("questionnaires").AutoMigrate(&types.Questionnaire{}); err != nil {
		return err
	}
	if err := a.db.Table("questions").AutoMigrate(&types.Question{}); err != nil {
		return err
	}
	if err := a.db.Table("options").AutoMigrate(&types.Option{}); err != nil {
		return err
	}
	if err := a.db.Table("response").AutoMigrate(&types.Response{}); err != nil {
		return err
	}
	return nil
}
