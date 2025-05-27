package dao

import (
	"fmt"
)

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"size:32"`
	Articles []Article // 用户拥有文章列表
}

type Article struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"size:32"`
	UserID uint   // 外键字段
	User   User   // 多对一
}

// AutoMigrate 自动迁移表结构
func AutoMigrate() error {
	if err := DB.AutoMigrate(&User{}, &Article{}); err != nil {
		return fmt.Errorf("auto migrate failed: %w", err)
	}
	return nil
}
