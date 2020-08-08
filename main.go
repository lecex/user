package main

import (
	// 公共引入
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	_ "github.com/micro/go-plugins/registry/nats/v2"

	// 执行数据迁移
	"github.com/lecex/user/config"
	"github.com/lecex/user/handler"
	_ "github.com/lecex/user/providers/migrations"
)

func main() {
	var Conf = config.Conf
	service := micro.NewService(
		micro.Name(Conf.Name),
		micro.Version(Conf.Version),
	)
	service.Init()
	// 注册服务
	h := handler.Handler{
		Server: service.Server(),
	}
	h.Register()
	// Run the server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	log.Log("serviser run ...")
}
