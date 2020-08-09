package handler

import (
	server "github.com/micro/go-micro/v2/server"

	db "github.com/lecex/user/providers/database"

	authPB "github.com/lecex/user/proto/auth"
	casbinPB "github.com/lecex/user/proto/casbin"
	frontPermitPB "github.com/lecex/user/proto/frontPermit"
	permissionPB "github.com/lecex/user/proto/permission"
	rolePB "github.com/lecex/user/proto/role"
	userPB "github.com/lecex/user/proto/user"

	"github.com/lecex/user/providers/casbin"
	"github.com/lecex/user/service"
	"github.com/lecex/user/service/repository"
)

// Register 注册
func Register(Server server.Server) {
	userPB.RegisterUsersHandler(Server, &User{&repository.UserRepository{db.DB}})                             // 用户服务实现
	authPB.RegisterAuthHandler(Server, &Auth{&service.TokenService{}, &repository.UserRepository{db.DB}})     // token 服务实现
	frontPermitPB.RegisterFrontPermitsHandler(Server, &FrontPermit{&repository.FrontPermitRepository{db.DB}}) // 前端权限服务实现
	permissionPB.RegisterPermissionsHandler(Server, &Permission{&repository.PermissionRepository{db.DB}})     // 权限服务实现
	rolePB.RegisterRolesHandler(Server, &Role{&repository.RoleRepository{db.DB}})                             // 角色服务实现
	casbinPB.RegisterCasbinHandler(Server, &Casbin{casbin.Enforcer})                                          // 权限管理服务实现
}
