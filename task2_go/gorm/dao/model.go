package dao

import (
	"fmt"
)

type User struct {
	ID       uint     `gorm:"primaryKey"`
	Name     string   `gorm:"size:32"`
	Age      int      `gorm:"default:0"`
	UserInfo UserInfo // 通过 UserInfo 可以拿到用户详细信息
}

type UserInfo struct {
	ID     uint   `gorm:"primaryKey"`
	Addr   string `gorm:"size:255"`
	Like   string `gorm:"size:255"`
	UserID uint   `gorm:"foreignKey:UserID"`
}

// AutoMigrate 自动迁移表结构
func AutoMigrate() error {
	if err := DB.AutoMigrate(&User{}, &UserInfo{}); err != nil {
		return fmt.Errorf("auto migrate failed: %w", err)
	}
	return nil
}
