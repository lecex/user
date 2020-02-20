package main

import (
	// 公共引入
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	// 执行数据迁移
	"github.com/lecex/user/config"
	"github.com/lecex/user/handler"
	_ "github.com/lecex/user/providers/migrations"
)

func main() {
	var Conf = config.Conf
	service := micro.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	service.Init()
	// 注册服务
	handler.Register(service.Server())
	// Run the server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	log.Log("serviser run ...")
}
