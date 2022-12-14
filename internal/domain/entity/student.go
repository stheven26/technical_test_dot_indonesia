package entity

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        int64          `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"column:name; type:VARCHAR(100)" json:"name"`
	Age       int            `gorm:"column:age; type:uint" json:"age"`
	School    string         `gorm:"column:school; type:VARCHAR(100)" json:"school"`
	Grade     int            `gorm:"column:grade" json:"grade"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at; index" json:"deleted_at"`
}
