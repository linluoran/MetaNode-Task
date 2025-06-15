package svc

import (
	"context"
	"fmt"
	"github.com/linluoran/common_rpc/usercenterrpcclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"usercenter/internal/config"
	"usercenter/internal/middleware"
	"usercenter/model"
)

type ServiceContext struct {
	Config               config.Config
	UserCreateMiddleware rest.Middleware
	UserModel            model.UserModel
	UserDataModel        model.UserDataModel
	UserRpcClient        usercenterrpcclient.UsercenterRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		UserCreateMiddleware: middleware.NewUserCreateMiddleware().Handle,
		UserModel: model.NewUserModel(
			sqlx.NewMysql(c.MySQL.DSN),
			c.Cache,
		),

		UserDataModel: model.NewUserDataModel(
			sqlx.NewMysql(c.MySQL.DSN),
			c.Cache,
		),

		UserRpcClient: usercenterrpcclient.NewUsercenterRpc(
			zrpc.MustNewClient(
				c.UserRpcConf,
				// 添加拦截器
				zrpc.WithUnaryClientInterceptor(TestClinetInterceptor),
			),
		),
	}
}

// 定义拦截器
func TestClinetInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("发送前")

	// 如果不写处理后逻辑 直接 return handler(ctx, req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	fmt.Println("发送后")
	return nil

}
