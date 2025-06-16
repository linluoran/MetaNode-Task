package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	GlobalConfig *BaseConfig
	configPath   = "./config/config.yaml"
)

type BaseConfig struct {
	Name  string      `mapstructure:"Name"`
	Env   string      `mapstructure:"Env"`
	Host  string      `mapstructure:"Host"`
	Port  string      `mapstructure:"Port"`
	Mysql MysqlConfig `mapstructure:"Mysql"`
}

type MysqlConfig struct {
	UserName        string `mapstructure:"Username"`
	Password        string `mapstructure:"Password"`
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"Host"`
	DBName          string `mapstructure:"DBName"`
	Timeout         string `mapstructure:"Timeout"`
	DSN             string `mapstructure:"DSN"`
	MaxOpenConns    int    `mapstructure:"MaxOpenConns"`
	MaxIdleConns    int    `mapstructure:"MaxIdleConns"`
	ConnMaxLifetime string `mapstructure:"ConnMaxLifetime"`
}

func LoadConfig() {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 自动绑定环境变量
	// viper.AutomaticEnv()

	// 读配置
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 实时监控配置变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		// 热更新配置
		if err := v.Unmarshal(Global); err != nil {
			panic(err)
		}
	})
}
