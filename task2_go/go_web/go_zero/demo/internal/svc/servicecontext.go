package svc

import (
	"demo/internal/config"
	"demo/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		/*
			添加 NewUserModel 即数据库表的模型 opts 为可选参数 忽略
				NewUserModel: 需要传入参数 conn sqlx.SqlConn, c cache.CacheConf
				conn 即 sqlx.NewMysql: 需要 datasource string
				是否需要传递缓存参数取决于, 生成Model时 cache 是否为 true
			所有的配置需要从 config.Config 里面获取
				先到 etc/demo-api.yaml 中添加数据库配置
		*/
		UserModel: model.NewUserModel(sqlx.NewMysql(c.MySQL.DSN)),
	}
}
