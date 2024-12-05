package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnectionConfig struct {
	Host   string
	Port   uint
	User   string
	Pass   string
	Dbname string
}

func (c DBConnectionConfig) MySqlDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.User, c.Pass, c.Host, c.Port, c.Dbname)
}

func NewMySqlGormConnection(config DBConnectionConfig) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.MySqlDSN()), &gorm.Config{
		Logger: logger.Discard,
	})
}
