package gorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Config 数据库默认配置
type Config struct {
	// Driver 主机连接方式
	Driver string
	// Host 主机地址
	Host string
	// Port 主机端口
	Port string
	// User 用户名
	User string
	// Password 密码
	Password string
	// DbName 数据库名称
	DbName string
	// Charset 数据库编码
	Charset string
}

// Connection 根据驱动创建连接
func Connection(conf *Config) (db *gorm.DB, err error) {
	if conf.Driver == "mysql" {
		return mysqlConnection(conf)
	}
	if conf.Driver == "postgres" {
		return postgresConnection(conf)
	}
	return db, fmt.Errorf(" '%v' driver doesn't exist. ", conf.Driver)
}
