package client

import (
	"context"

	"github.com/micro/go-micro/v2/client"
	cli "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/config/cmd"
)

func init() {
	cmd.Init()
}

// Call 使用默认客户端对服务进行同步调用
func Call(ctx context.Context, service string, endpoint string, req interface{}, rsp interface{}, opts ...client.CallOption) error {
	client.DefaultClient = cli.NewClient()
	request := client.NewRequest(service, endpoint, req)
	err := client.Call(ctx, request, rsp, opts...)
	return err
}
