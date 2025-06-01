package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Student struct {
	ID    uint   `gorm:"primary_key; auto_increment"`
	Name  string `gorm:"size(32)"`
	Age   uint
	Grade string `gorm:"size(16)"`
}

type Account struct {
	ID      uint    `gorm:"primary_key; auto_increment"`
	Balance float64 `gorm:"type:decimal(10,2)"`
}

type Transaction struct {
	ID          uint `gorm:"primary_key; auto_increment"`
	FromAccount uint `gorm:"unique_index:idx_from_account"`
	ToAccount   uint `gorm:"unique_index:idx_to_account"`
}

// 仅通过ORM建表
type employee struct {
	ID         uint    `gorm:"primary_key; auto_increment"`
	Name       string  `gorm:"size(16)"`
	Department string  `gorm:"size(16)"`
	Salary     float64 `gorm:"type:decimal(10,2)"`
}

// 仅通过ORM建表
type book struct {
	ID     uint    `gorm:"primary_key; auto_increment"`
	Title  string  `gorm:"size(16)"`
	Author string  `gorm:"size(16)"`
	Price  float64 `gorm:"type:decimal(10,2)"`
}

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size(16)"`
	Posts []Post
}

type Post struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:16"`
	WordCount     int
	User          User
	UserID        uint
	CommentStatus string `gorm:"size(8)"`
	Comments      []Comment
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"size:255"`
	Post    Post
	PostID  uint
}

// BeforeCreate  创建文章时自动写入字数
func (p *Post) BeforeCreate(db *gorm.DB) error {
	if p.WordCount != 0 {
		// 传入结构体已经设置了字数
		return nil
	}

	// 获取参数
	tmpWC := db.Statement.Context.Value("WordCount")
	value, isInt := tmpWC.(int)
	if !isInt {
		log.Printf("BeforeCreate: 字数不是整数格式")
		return errors.New("字数不是整数格式")
	}
	p.WordCount = value
	return nil
}

func (c *Comment) AfterDelete(db *gorm.DB) error {

	if c.PostID == 0 {
		// 对应文章不存在
		log.Printf("AfterDelete: 对应文章不存在")
		return nil
	}

	var commentCount int64
	db.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&commentCount)

	if commentCount == 0 {
		if upErr := db.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("CommentStatus", "无评论").
			Error; upErr != nil {
			log.Println(upErr)
			return upErr
		}
	}
	return nil
}

// AutoMigrate 自动迁移表结构
func AutoMigrate() error {
	if err := DB.AutoMigrate(
		&Student{},
		&Account{},
		&Transaction{},
		&employee{},
		&book{},
		&User{},
		&Post{},
		&Comment{},
	); err != nil {
		return fmt.Errorf("auto migrate failed: %w", err)
	}
	return nil
}
