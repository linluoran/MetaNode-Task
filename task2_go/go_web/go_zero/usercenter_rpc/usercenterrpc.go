package main

import (
	"context"
	"flag"
	"fmt"

	"usercenter_rpc/internal/config"
	"usercenter_rpc/internal/server"
	"usercenter_rpc/internal/svc"
	"usercenter_rpc/usercenter_rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/usercenterrpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		usercenter_rpc.RegisterUsercenterRpcServer(grpcServer, server.NewUsercenterRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 添加拦截器
	s.AddUnaryInterceptors(TestServerInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

// 定义拦截器
func TestServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println("TestServerInterceptor start")

	fmt.Printf("TestServerInterceptor req %+v\n", req)
	fmt.Printf("TestServerInterceptor info %+v\n", info)

	reps, err := handler(ctx, req)

	fmt.Println("TestServerInterceptor end")
	return reps, err
}
