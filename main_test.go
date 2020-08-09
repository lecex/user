package main

import (
	"context"
	"testing"

	_ "github.com/lecex/user/providers/migrations"
	"github.com/lecex/user/service"

	handler "github.com/lecex/user/handler"
	authPB "github.com/lecex/user/proto/auth"
	frontPermitPB "github.com/lecex/user/proto/frontPermit"
	permissionPB "github.com/lecex/user/proto/permission"
	userPB "github.com/lecex/user/proto/user"
	db "github.com/lecex/user/providers/database"

	"github.com/lecex/user/service/repository"
)

func TestFrontPermitUpdateOrCreate(t *testing.T) {
	req := &frontPermitPB.Request{
		FrontPermit: &frontPermitPB.FrontPermit{
			App: "ui", Service: "role", Method: "permission", Name: "角色权限", Description: "管理角色权限。",
		},
	}
	res := &frontPermitPB.Response{}
	h := handler.FrontPermit{&repository.FrontPermitRepository{db.DB}}
	err := h.UpdateOrCreate(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
func TestPermissionsSync(t *testing.T) {
	req := &permissionPB.Request{
		Permissions: []*permissionPB.Permission{
			{Service: "Users", Method: "Create", Auth: true, Policy: true, Name: "创建用户", Description: "创建新用户权限。"},
			{Service: "Users", Method: "Exist", Auth: false, Policy: false, Name: "检测用户", Description: "检测用户是否存在权限。"},
			{Service: "Users", Method: "Get", Auth: true, Policy: true, Name: "用户查询", Description: "查询用户信息权限。"},
			{Service: "Auth", Method: "Auth", Auth: false, Policy: false, Name: "用户授权", Description: "用户登录授权返回 token 权限。"},
			{Service: "Auth", Method: "ValidateToken", Auth: false, Policy: false, Name: "权限认证", Description: "权限相关认证权限。"},
		},
	}
	res := &permissionPB.Response{}
	h := handler.Permission{&repository.PermissionRepository{db.DB}}
	err := h.Sync(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
func TestPermissionsUpdateOrCreate(t *testing.T) {
	req := &permissionPB.Request{
		Permission: &permissionPB.Permission{
			Service: "user-api", Method: "Auth.Auth1", Name: "用户授权3", Description: "用户登录授权返回 token 权限。",
		},
	}
	res := &permissionPB.Response{}
	h := handler.Permission{&repository.PermissionRepository{db.DB}}
	err := h.UpdateOrCreate(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
func TestUserCreate(t *testing.T) {
	req := &userPB.Request{
		User: &userPB.User{
			Username: `bvbv0115`,
			Password: `123456`,
			Mobile:   `13953186115`,
			Email:    `bvbv0a115@qq.com`,
			Name:     `bvbv0111`,
			Avatar:   `https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif`,
			Origin:   `user`,
		},
	}
	res := &userPB.Response{}
	h := handler.User{&repository.UserRepository{db.DB}}
	err := h.Create(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestUserIsExist(t *testing.T) {
	req := &userPB.Request{
		User: &userPB.User{
			Username: `bvbv0115`,
			Mobile:   `13953186115`,
			Email:    `bvbv0a115@qq.com`,
		},
	}
	res := &userPB.Response{}
	h := handler.User{&repository.UserRepository{db.DB}}
	err := h.Exist(context.TODO(), req, res)
	// fmt.Println("exist", req, res.Valid, err)
	t.Log(req, res, err)
}
func TestUserGet(t *testing.T) {
	req := &userPB.Request{
		User: &userPB.User{
			Username: `bvbv0115`,
		},
	}
	res := &userPB.Response{}
	h := handler.User{&repository.UserRepository{db.DB}}
	err := h.Get(context.TODO(), req, res)
	// fmt.Println("UserGet", res, err)
	t.Log(req, res, err)
}

func TestUserList(t *testing.T) {
	req := &userPB.Request{
		ListQuery: &userPB.ListQuery{
			Limit: 20,
			Page:  1,
			Sort:  "created_at desc",
		},
	}
	res := &userPB.Response{}
	h := handler.User{&repository.UserRepository{db.DB}}
	err := h.List(context.TODO(), req, res)
	// fmt.Println("UserList", res, err)
	t.Log(req, res, err)
}
func TestUserUpdate(t *testing.T) {
	// req := &userPB.Request{
	// 	User: &userPB.User{
	// 		Id:       `8cd1d57f-6f53-49e4-b751-96eefc4f4b20`,
	// 		Username: `bvbv0111`,
	// 		Name:     `newbvbv`,
	// 	},
	// }
	// res := &userPB.Response{}
	// handler := &handler.Handler{}
	// h := handler.User()
	// err := h.Update(context.TODO(), req, res)
	// fmt.Println("UserUpdate", req, res, err)
	// t.Log(req, res, err)
}
func TestUserDelete(t *testing.T) {
	req := &userPB.Request{
		User: &userPB.User{
			Id: `8cd1d57f-6f53-49e4-b751-96eefc4f4b20`,
		},
	}
	res := &userPB.Response{}
	h := handler.User{&repository.UserRepository{db.DB}}
	err := h.Delete(context.TODO(), req, res)
	// fmt.Println("UserDelete", req, res, err)
	t.Log(req, res, err)
}

// Auth
func TestAuth(t *testing.T) {
	req := &authPB.Request{
		User: &authPB.User{
			Username: `bvbv011`,
			Password: `123456`,
		},
	}
	res := &authPB.Response{}
	h := handler.Auth{&service.TokenService{}, &repository.UserRepository{db.DB}}
	err := h.Auth(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestAuthById(t *testing.T) {
	req := &authPB.Request{
		User: &authPB.User{
			Id: `c0a83b24-c01c-4601-a1c6-17e3c1864c5a`,
		},
	}
	res := &authPB.Response{}
	h := handler.Auth{&service.TokenService{}, &repository.UserRepository{db.DB}}
	err := h.AuthById(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestValidateToken(t *testing.T) {
	req := &authPB.Request{
		Token: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiYzBhODNiMjQtYzAxYy00NjAxLWExYzYtMTdlM2MxODY0YzVhIn0sImV4cCI6MTU3MDM0OTc2MCwiaXNzIjoidXNlciJ9.Y3l55bE3StZL66RPbrTk8zVgUZBll0Pc6yV6ljb22k4`,
	}
	res := &authPB.Response{}
	h := handler.Auth{&service.TokenService{}, &repository.UserRepository{db.DB}}
	err := h.ValidateToken(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
