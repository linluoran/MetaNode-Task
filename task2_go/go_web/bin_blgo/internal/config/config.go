package config

import (
	"github.com/spf13/viper"
)

var (
	GlobalConfig *BaseConfig
	configPath   = "etc/config.yaml"
)

type BaseConfig struct {
	Name  string       `mapstructure:"Name"`
	Env   string       `mapstructure:"Env"`
	Host  string       `mapstructure:"Host"`
	Port  string       `mapstructure:"Port"`
	Mysql MysqlConfig  `mapstructure:"Mysql"`
	Log   LoggerConfig `mapstructure:"Log"`
}

type MysqlConfig struct {
	UserName        string `mapstructure:"Username"`
	Password        string `mapstructure:"Password"`
	Host            string `mapstructure:"Host"`
	Port            string `mapstructure:"Port"`
	DBName          string `mapstructure:"DBName"`
	Timeout         string `mapstructure:"Timeout"`
	DSN             string `mapstructure:"DSN"`
	MaxOpenConns    int    `mapstructure:"MaxOpenConns"`
	MaxIdleConns    int    `mapstructure:"MaxIdleConns"`
	ConnMaxLifetime string `mapstructure:"ConnMaxLifetime"`
}

type LoggerConfig struct {
	LogPath    string `mapstructure:"LogPath"`
	MaxSize    int    `mapstructure:"MaxSize"`
	MaxBackups int    `mapstructure:"MaxBackups"`
	MaxAge     int    `mapstructure:"MaxAge"`
	Compress   bool   `mapstructure:"Compress"`
}

func LoadConfig() error {
	GlobalConfig = &BaseConfig{}
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 自动绑定环境变量
	// viper.AutomaticEnv()

	// 读配置
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		return err
	}
	return nil
}
