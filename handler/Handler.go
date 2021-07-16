package handler

import (
	db "github.com/lecex/user/providers/database"

	authPB "github.com/lecex/user/proto/auth"
	casbinPB "github.com/lecex/user/proto/casbin"
	frontPermitPB "github.com/lecex/user/proto/frontPermit"
	permissionPB "github.com/lecex/user/proto/permission"
	rolePB "github.com/lecex/user/proto/role"
	secretKeyPB "github.com/lecex/user/proto/secretKey"
	userPB "github.com/lecex/user/proto/user"

	"github.com/lecex/user/providers/casbin"
	"github.com/lecex/user/service"
	"github.com/lecex/user/service/repository"
	"github.com/micro/go-micro/v2"
)

const topic = "event"

// Register 注册
func Register(srv micro.Service) {
	server := srv.Server()
	publisher := micro.NewPublisher(topic, srv.Client())
	userPB.RegisterUsersHandler(server, &User{&repository.UserRepository{db.DB}, publisher})
	secretKeyPB.RegisterSecretKeysHandler(server, &SecretKey{&repository.SecretKeyRepository{db.DB}})         // 用户服务实现
	authPB.RegisterAuthHandler(server, &Auth{&service.TokenService{}, &repository.UserRepository{db.DB}})     // token 服务实现
	frontPermitPB.RegisterFrontPermitsHandler(server, &FrontPermit{&repository.FrontPermitRepository{db.DB}}) // 前端权限服务实现
	permissionPB.RegisterPermissionsHandler(server, &Permission{&repository.PermissionRepository{db.DB}})     // 权限服务实现
	rolePB.RegisterRolesHandler(server, &Role{&repository.RoleRepository{db.DB}})                             // 角色服务实现
	casbinPB.RegisterCasbinHandler(server, &Casbin{casbin.Enforcer})                                          // 权限管理服务实现
}
