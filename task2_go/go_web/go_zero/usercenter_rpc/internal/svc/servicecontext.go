package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"usercenter_rpc/internal/config"
	"usercenter_rpc/model"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	UserDataModel model.UserDataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(
			sqlx.NewMysql(c.MySQL.DSN),
			c.Cache,
		),

		UserDataModel: model.NewUserDataModel(
			sqlx.NewMysql(c.MySQL.DSN),
			c.Cache,
		),
	}
}
