package config

import (
	"github.com/spf13/viper"
	"time"
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
	Jwt   JwtConfig    `mapstructure:"Jwt"`
}

type (
	MysqlConfig struct {
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

	LoggerConfig struct {
		LogPath    string `mapstructure:"LogPath"`
		MaxSize    int    `mapstructure:"MaxSize"`
		MaxBackups int    `mapstructure:"MaxBackups"`
		MaxAge     int    `mapstructure:"MaxAge"`
		Compress   bool   `mapstructure:"Compress"`
	}

	JwtConfig struct {
		Secret      string        `mapstructure:"Secret"`
		TokenExpire time.Duration `mapstructure:"TokenExpire"`
		Issuer      string        `mapstructure:"Issuer"`
	}
)

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
