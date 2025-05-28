package dao

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Info struct {
	Status string `json:"status"`
	Addr   string `json:"addr"`
	Age    int    `json:"age"`
}

type AuthModel struct {
	ID   uint64 `gorm:"primary_key"`
	Name string
	Info Info `gorm:"type:string;"`
}

// Scan 从数据库读出来
func (i *Info) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Failed to umarshal JSON value: %s", value))
	}
	err := json.Unmarshal(bytes, i)
	return err
}

func (i Info) Value() (driver.Value, error) {
	return json.Marshal(i)
}

// AutoMigrate 自动迁移表结构
func AutoMigrate() error {
	if err := DB.AutoMigrate(
		&AuthModel{},
	); err != nil {
		return fmt.Errorf("auto migrate failed: %w", err)
	}
	return nil
}
