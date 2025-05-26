// dao/load_conf.go

package dao

import (
	"github.com/spf13/viper"
)

type MySQLConfig struct {
	UserName        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	DBName          string `mapstructure:"dbname"`
	Timeout         string `mapstructure:"timeout"`
	DSN             string `mapstructure:"dsn"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime string `mapstructure:"conn_max_lifetime"`
}

// LoadConfig 从指定路径加载 MySQL 配置
// 参数:
//   - path: 配置文件路径
//
// 返回值:
//   - *MySQLConfig: 成功时返回解析后的 MySQL 配置对象指针
//   - error: 失败时返回错误信息
func LoadConfig(path string) (*MySQLConfig, error) {
	// 设置 Viper 的配置文件路径
	viper.SetConfigFile(path)

	// 启用自动绑定环境变量功能（环境变量名将自动匹配配置键名）
	// viper.AutomaticEnv()

	// 读取并解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// 创建 MySQLConfig 结构体实例
	var cfg MySQLConfig
	// 从配置文件中解析 "mysql" 段落到 cfg 结构体
	if err := viper.UnmarshalKey("mysql", &cfg); err != nil {
		return nil, err
	}

	// 返回解析后的配置对象
	return &cfg, nil
}
