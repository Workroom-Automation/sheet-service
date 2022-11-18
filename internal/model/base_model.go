package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
	CreatedBy string         `gorm:"column:created_by" json:"created_by"`
	UpdatedBy string         `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy string         `gorm:"column:deleted_by" json:"deleted_by"`
	Namespace string         `gorm:"column:namespace" json:"namespace"`
}
