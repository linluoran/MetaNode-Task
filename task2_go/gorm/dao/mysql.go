package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

func InitMySQL(cfg *MySQLConfig) error {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Timeout)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用事务
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// 配置连接池
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	duration, _ := time.ParseDuration(cfg.ConnMaxLifetime)
	sqlDB.SetConnMaxLifetime(duration)

	DB = db
	return nil
}

func init() {
	// 加载配置
	cfg, loadErr := LoadConfig("./config/config.yaml")
	if loadErr != nil {
		log.Fatalf("Load config failed: %v", loadErr)
	}

	// 初始化数据库
	if err := InitMySQL(cfg); err != nil {
		log.Fatalf("MySQL init failed: %v", err)
	}

	setErr := DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
	if setErr != nil {
		log.Fatalf("Failed to setup join table: %v", setErr)
	}

	// 自动迁移 不是每次都需要迁移
	//if err := AutoMigrate(); err != nil {
	//	log.Fatalf("Auto migrate failed: %v", err)
	//}

}
