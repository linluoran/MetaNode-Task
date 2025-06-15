package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// 加载 MySQL 配置
	MySQL struct{ DSN string }

	// 加载 Cache 配置
	Cache cache.CacheConf
}
