// Package models 模型通用属性和方法
package models

import "time"

// BaseModel 模型积累
type BaseModel struct {
	ID uint64 `gorm:"colum:id;primaryKey;autoIncrement;"json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"colum:created_at;index"json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"colum:updated_at;index;json:"updated_at,omitempty"`
}
