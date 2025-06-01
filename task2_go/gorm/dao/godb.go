package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 匿名导入驱动
	"log"
	"time"
)

var GoDB *sql.DB

func InitDB(cfg *MySQLConfig) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Timeout)

	db, openErr := sql.Open("mysql", dsn)
	if openErr != nil {
		log.Fatal(openErr)
		return openErr
	}

	db.SetMaxOpenConns(25)                  // 最大并发连接数
	db.SetMaxIdleConns(10)                  // 空闲连接保留数
	db.SetConnMaxLifetime(30 * time.Minute) // 连接最大存活时间

	// 验证连接
	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal("连接失败:", pingErr)
		return pingErr
	}

	// 获取连接池实时统计信息
	stats := db.Stats()
	fmt.Printf("当前连接数: %d (空闲: %d, 使用中: %d)\n",
		stats.OpenConnections,
		stats.Idle,
		stats.InUse)

	GoDB = db
	return nil
}

func init() {
	// 加载配置
	cfg, loadErr := LoadConfig("./config/config.yaml")
	if loadErr != nil {
		log.Fatalf("Load config failed: %v", loadErr)
	}

	// 初始化数据库
	if err := InitDB(cfg); err != nil {
		log.Fatalf("MySQL init failed: %v", err)
	}
}
