package entity

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ID            int64  `gorm:"primarykey" json:"id"`
	StudentID     int64  `gorm:"student_id" json:"studentId"`
	Subject       string `gorm:"subject; type:VARCHAR(50)" json:"subject"`
	DurationClass int    `gorm:"duration_class" json:"durationClass"`
	//Relations
	Student   *Student       `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at; index" json:"deleted_at"`
}
