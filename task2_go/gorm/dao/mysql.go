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

	//// 自动迁移 不是每次都需要迁移
	//if err := AutoMigrate(); err != nil {
	//	log.Fatalf("Auto migrate failed: %v", err)
	//}

	var users []User
	DB.Find(&users).Unscoped().Delete(&User{})
	users = []User{
		{ID: 1, Name: "李元芳", Age: 32, Email: PtrString("lyf@yf.com"), Gender: true},
		{ID: 2, Name: "张武", Age: 18, Email: PtrString("zhangwu@lly.cn"), Gender: true},
		{ID: 3, Name: "极枫", Age: 23, Email: PtrString("ff@yahoo.com"), Gender: true},
		{ID: 4, Name: "刘大", Age: 54, Email: PtrString("Liuda@qq.com"), Gender: true},
		{ID: 5, Name: "李武", Age: 23, Email: PtrString("liwu@lly.cn"), Gender: true},
		{ID: 6, Name: "李琦", Age: 14, Email: PtrString("liqi@lly.cn"), Gender: false},
		{ID: 7, Name: "晓梅", Age: 25, Email: PtrString("xiaomeo@sl.com"), Gender: false},
		{ID: 8, Name: "如燕", Age: 26, Email: PtrString("ruyan@yf.com"), Gender: false},
		{ID: 9, Name: "魔灵", Age: 21, Email: PtrString("moling@sl.com"), Gender: true},
	}
	DB.Create(&users)

}

func PtrString(s string) *string {
	return &s
}
