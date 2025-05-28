package dao

import (
	"fmt"
)

type Tag struct {
	ID   uint `gorm:"primary_key"`
	Name string
	// article_tags 指定第三张表的名字
	Articles []Article `gorm:"many2many:article_tags;"`
}
type Article struct {
	ID    uint `gorm:"primary_key"`
	Title string
	Tags  []Tag `gorm:"many2many:article_tags;"`
}

// AutoMigrate 自动迁移表结构
func AutoMigrate() error {
	if err := DB.AutoMigrate(&Tag{}, &Article{}); err != nil {
		return fmt.Errorf("auto migrate failed: %w", err)
	}
	return nil
}
