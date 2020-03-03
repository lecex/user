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

// Handler 注册方法
type Handler struct {
	Server server.Server
}

// Register 注册
func (srv *Handler) Register() {
	userPB.RegisterUsersHandler(srv.Server, srv.User())                      // 用户服务实现
	authPB.RegisterAuthHandler(srv.Server, srv.Auth())                       // token 服务实现
	frontPermitPB.RegisterFrontPermitsHandler(srv.Server, srv.FrontPermit()) // 前端权限服务实现
	permissionPB.RegisterPermissionsHandler(srv.Server, srv.Permission())    // 权限服务实现
	rolePB.RegisterRolesHandler(srv.Server, srv.Role())                      // 角色服务实现
	casbinPB.RegisterCasbinHandler(srv.Server, srv.Casbin())                 // 权限管理服务实现
}

// User 用户管理服务实现
func (srv *Handler) User() *User {
	return &User{&repository.UserRepository{db.DB}}
}

// Auth 授权管理服务实现
func (srv *Handler) Auth() *Auth {
	return &Auth{&service.TokenService{}, &repository.UserRepository{db.DB}}
}

// FrontPermit 前端权限管理服务实现
func (srv *Handler) FrontPermit() *FrontPermit {
	return &FrontPermit{&repository.FrontPermitRepository{db.DB}}
}

// Permission 权限管理服务实现
func (srv *Handler) Permission() *Permission {
	return &Permission{&repository.PermissionRepository{db.DB}}
}

// Role 角色管理服务实现
func (srv *Handler) Role() *Role {
	return &Role{&repository.RoleRepository{db.DB}}
}

// Casbin 权限管理服务实现
func (srv *Handler) Casbin() *Casbin {
	return &Casbin{casbin.Enforcer}
}
