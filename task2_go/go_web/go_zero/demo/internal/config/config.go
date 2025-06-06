package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	// 加载 MySQL 配置
	MySQL struct {
		DSN string
	}
}
