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
type Handler struct {}

// Register 注册
func Register(s server.Server)  {
	h := &Handler{}
	userPB.RegisterUsersHandler(s, h.User())	// 用户服务实现
	authPB.RegisterAuthHandler(s, h.Auth())	// token 服务实现
	frontPermitPB.RegisterFrontPermitsHandler(s, h.FrontPermit())	// 前端权限服务实现
	permissionPB.RegisterPermissionsHandler(s, h.Permission())	// 权限服务实现
	rolePB.RegisterRolesHandler(s, h.Role())	// 角色服务实现
	casbinPB.RegisterCasbinHandler(s, h.Casbin()) // 权限管理服务实现
}

// User 用户管理服务实现
func (srv *Handler) User() *User {
	repo := &repository.UserRepository{db.DB}
	return  &User{repo}
}

// Auth 授权管理服务实现
func (srv *Handler) Auth() *Auth {
	token := &service.TokenService{}
	repo := &repository.UserRepository{db.DB}
	return  &Auth{token, repo}
}

// FrontPermit 前端权限管理服务实现
func (srv *Handler) FrontPermit() *FrontPermit {
	repo := &repository.FrontPermitRepository{db.DB}
	return  &FrontPermit{repo}
}

// Permission 权限管理服务实现
func (srv *Handler) Permission() *Permission {
	repo := &repository.PermissionRepository{db.DB}
	return  &Permission{repo}
}

// Role 角色管理服务实现
func (srv *Handler) Role() *Role {
	repo := &repository.RoleRepository{db.DB}
	return  &Role{repo}
}

// Casbin 权限管理服务实现
func (srv *Handler) Casbin() *Casbin {
	return   &Casbin{casbin.Enforcer}
}



