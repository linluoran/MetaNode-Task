package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	// 加载 MySQL 配置
	MySQL struct{ DSN string }

	// 加载 Cache 配置
	Cache cache.CacheConf
}
