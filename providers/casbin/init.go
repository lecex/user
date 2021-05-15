package casbin

import (
	"fmt"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"

	"github.com/lecex/core/env"
)

// DB 管理包
var (
	// DB 连接
	Enforcer *casbin.Enforcer
)

func init() {
	Driver := env.Getenv("DB_DRIVER", "mysql")
	// Host 主机地址
	Host := env.Getenv("DB_HOST", "127.0.0.1")
	// Port 主机端口
	Port := env.Getenv("DB_PORT", "3306")
	// User 用户名
	User := env.Getenv("DB_USER", "root")
	// Password 密码
	Password := env.Getenv("DB_PASSWORD", "123456")
	// DbName 数据库名称
	DbName := env.Getenv("DB_NAME", "user")
	// Charset 数据库编码
	Charset := env.Getenv("DB_CHARSET", "utf8")
	// Initialize the model from a string.
	m := casbin.NewModel(`
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || g(r.sub,"root")
	`)
	a := gormadapter.NewAdapter(Driver, fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		User, Password, Host, Port, DbName, Charset,
	), true) // Your driver and data source.
	Enforcer = casbin.NewEnforcer(m, a)
	Enforcer.LoadPolicy()
}
