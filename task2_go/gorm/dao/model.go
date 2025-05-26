package dao

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID     uint   `gorm:"primarykey"`
	Name   string `gorm:"size:32"`
	Age    uint8
	Gender bool
	Email  *string `gorm:"size:64"`
}

// AutoMigrate 自动迁移表结构
func AutoMigrate() error {
	if err := DB.AutoMigrate(&User{}); err != nil {
		return fmt.Errorf("auto migrate failed: %w", err)
	}
	return nil
}
