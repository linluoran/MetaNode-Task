package svc

import (
	"github.com/linluoran/common_rpc/usercenterrpcclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
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

		UserRpcClient: usercenterrpcclient.NewUsercenterRpc(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
