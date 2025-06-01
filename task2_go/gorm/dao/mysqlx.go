package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var DBx *sqlx.DB

func InitDBx(cfg *MySQLConfig) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Timeout)

	// 使用 sql.Open 连接数据库
	dbx, openErr := sqlx.Open("mysql", dsn)
	if openErr != nil {
		log.Fatal(openErr)
		return openErr
	}

	dbx.SetMaxOpenConns(25)                  // 最大并发连接数
	dbx.SetMaxIdleConns(10)                  // 空闲连接保留数
	dbx.SetConnMaxLifetime(30 * time.Minute) // 连接最大存活时间

	// 验证连接
	if pingErr := dbx.Ping(); pingErr != nil {
		log.Fatal("连接失败:", pingErr)
		return pingErr
	}

	DBx = dbx
	return nil
}

func init() {
	cfg, loadErr := LoadConfig("./config/config.yaml")
	if loadErr != nil {
		log.Fatalf("Load config failed: %v", loadErr)
	}

	// 初始化数据库
	if err := InitDBx(cfg); err != nil {
		log.Fatalf("MySQL init failed: %v", err)
	}
}

type Userx struct {
	ID        int       `json:"id"`
	Name      string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
