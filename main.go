package main

import (
	// 公共引入
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"

	// 执行数据迁移
	_ "github.com/lecex/user/providers/migrations"
	db "github.com/lecex/user/providers/database"

	authPB "github.com/lecex/user/proto/auth"
	casbinPB "github.com/lecex/user/proto/casbin"
	frontPermitPB "github.com/lecex/user/proto/frontPermit"
	permissionPB "github.com/lecex/user/proto/permission"
	rolePB "github.com/lecex/user/proto/role"
	userPB "github.com/lecex/user/proto/user"

	"github.com/lecex/user/providers/casbin"
	"github.com/lecex/user/handler"
	"github.com/lecex/user/service"
	"github.com/lecex/user/service/repository"
)

func main() {
	config.LoadFile("config.yaml")

	srv := micro.NewService(
		micro.Name(config.Get("service", "name").String("user")),
		micro.Version(config.Get("service", "bersion").String("latest")),
	)
	srv.Init()

	// 用户服务实现
	repo := &repository.UserRepository{db.DB}
	userPB.RegisterUsersHandler(srv.Server(), &handler.User{repo})

	// token 服务实现
	token := &service.TokenService{}
	authPB.RegisterAuthHandler(srv.Server(), &handler.Auth{token, repo})

	// 前端权限服务实现
	fprepo := &repository.FrontPermitRepository{db.DB}
	frontPermitPB.RegisterFrontPermitsHandler(srv.Server(), &handler.FrontPermit{fprepo})

	// 权限服务实现
	prepo := &repository.PermissionRepository{db.DB}
	permissionPB.RegisterPermissionsHandler(srv.Server(), &handler.Permission{prepo})

	// 角色服务实现
	rrepo := &repository.RoleRepository{db.DB}
	rolePB.RegisterRolesHandler(srv.Server(), &handler.Role{rrepo})

	// 权限管理服务实现
	casbinPB.RegisterCasbinHandler(srv.Server(), &handler.Casbin{casbin.Enforcer})
	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
	log.Log("serviser run ...")
}
