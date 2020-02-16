package gorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// postgresConnection 创建数据库连接
func postgresConnection(conf *Config) (*gorm.DB, error) {
	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			conf.Host, conf.Port, conf.User, conf.DbName, conf.Password,
		),
	)
}
