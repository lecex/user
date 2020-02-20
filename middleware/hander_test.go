package middleware

import (
	"context"
	"fmt"
	"testing"

	PB "github.com/lecex/user/proto/permission"
	"github.com/micro/go-micro/v2/server"
)

func TestConfig(t *testing.T) {
	Pata := []*PB.Permission{
		{Service: "Users", Method: "Create", Auth: false, Policy: false, Name: "创建用户", Description: "创建新用户权限。"},
		{Service: "Users", Method: "Exist", Auth: false, Policy: false, Name: "检测用户", Description: "检测用户是否存在权限。"},
		{Service: "Users", Method: "Get", Auth: true, Policy: false, Name: "用户查询", Description: "查询用户信息权限。"},
		{Service: "Auth", Method: "Auth", Auth: false, Policy: false, Name: "用户授权", Description: "用户登录授权返回 token 权限。"},
		{Service: "Auth", Method: "ValidateToken", Auth: false, Policy: false, Name: "权限认证", Description: "权限相关认证权限。"},
	}
	h := Handler{Permissions: Pata}
	hf := h.Wrapper(HandlerFunc)
	err := hf(context.Background(), nil, nil)
	fmt.Println(err)
}
func HandlerFunc(ctx context.Context, req server.Request, rsp interface{}) error {
	fmt.Println(ctx, req, rsp)
	return nil
}
