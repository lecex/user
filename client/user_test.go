package client

import (
	"fmt"
	"testing"

	PB "github.com/lecex/user/proto/permission"
)

func TestConfig(t *testing.T) {
	Pata := []*PB.Permission{
		{Service: "Users", Method: "Create", Auth: true, Policy: true, Name: "创建用户", Description: "创建新用户权限。"},
		{Service: "Users", Method: "Exist", Auth: false, Policy: false, Name: "检测用户", Description: "检测用户是否存在权限。"},
		{Service: "Users", Method: "Get", Auth: true, Policy: true, Name: "用户查询", Description: "查询用户信息权限。"},
		{Service: "Auth", Method: "Auth", Auth: false, Policy: false, Name: "用户授权", Description: "用户登录授权返回 token 权限。"},
		{Service: "Auth", Method: "ValidateToken", Auth: false, Policy: false, Name: "权限认证", Description: "权限相关认证权限。"},
	}
	h := User{ServiceName: `user.srv`}
	err := h.SyncPermission(Pata)
	fmt.Println(err)
}
