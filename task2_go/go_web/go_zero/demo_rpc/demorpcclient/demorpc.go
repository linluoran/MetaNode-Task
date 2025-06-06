// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3
// Source: demo_rpc.proto

package demoRpcClient

import (
	"context"

	"demo_rpc/demo_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = demo_rpc.Request
	Response = demo_rpc.Response

	DemoRpc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultDemoRpc struct {
		cli zrpc.Client
	}
)

func NewDemoRpc(cli zrpc.Client) DemoRpc {
	return &defaultDemoRpc{
		cli: cli,
	}
}

func (m *defaultDemoRpc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := demo_rpc.NewDemoRpcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
