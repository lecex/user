package main

import (
	// 公共引入
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	// 执行数据迁移
	_ "github.com/lecex/user/database/migrations"
	db "github.com/lecex/user/providers/database"

	authPB "github.com/lecex/user/proto/auth"
	casbinPB "github.com/lecex/user/proto/casbin"
	frontPermitPB "github.com/lecex/user/proto/frontPermit"
	permissionPB "github.com/lecex/user/proto/permission"
	rolePB "github.com/lecex/user/proto/role"
	userPB "github.com/lecex/user/proto/user"

	"github.com/lecex/user/providers/casbin"
	"github.com/lecex/user/hander"
	"github.com/lecex/user/service"
)

func main() {
	service := micro.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	service.Init()

	// 用户服务实现
	repo := &service.UserRepository{db.DB}
	userPB.RegisterUsersHandler(srv.Server(), &hander.User{repo})

	// token 服务实现
	token := &service.TokenService{}
	authPB.RegisterAuthHandler(srv.Server(), &hander.Auth{token, repo})

	// 前端权限服务实现
	fprepo := &service.FrontPermitRepository{db.DB}
	frontPermitPB.RegisterFrontPermitsHandler(srv.Server(), &hander.FrontPermit{fprepo})

	// 权限服务实现
	prepo := &service.PermissionRepository{db.DB}
	permissionPB.RegisterPermissionsHandler(srv.Server(), &hander.Permission{prepo})

	// 角色服务实现
	rrepo := &service.RoleRepository{db.DB}
	rolePB.RegisterRolesHandler(srv.Server(), &hander.Role{rrepo})

	// 权限管理服务实现
	casbinPB.RegisterCasbinHandler(srv.Server(), &hander.Casbin{casbin.Enforcer})
	// Run the server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	log.Log("serviser run ...")
}
