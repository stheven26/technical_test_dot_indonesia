package repository

import (
	"context"
	"technical-test/internal/domain/entity"

	"gorm.io/gorm"
)

type classRepository struct {
	db *gorm.DB
}

type Class interface {
	Create(ctx context.Context, entity *entity.Class) (err error)
	GetAll(ctx context.Context, entity *[]entity.Class) (err error)
	GetById(ctx context.Context, id int64, entity *entity.Class) (err error)
	Delete(ctx context.Context, entity *entity.Class) (err error)
}

func NewClassRepository(db *gorm.DB) Class {
	if db == nil {
		panic("NewClassRepository is nil")
	}
	return &classRepository{db}
}

func (c *classRepository) Create(ctx context.Context, entity *entity.Class) (err error) {
	err = c.db.WithContext(ctx).Create(&entity).Error
	return
}

func (c *classRepository) GetAll(ctx context.Context, entity *[]entity.Class) (err error) {
	err = c.db.WithContext(ctx).Find(&entity).Error
	return
}

func (c *classRepository) GetById(ctx context.Context, id int64, entity *entity.Class) (err error) {
	err = c.db.WithContext(ctx).Where(`id=?`, id).Find(&entity).Error
	return
}

func (c *classRepository) Delete(ctx context.Context, entity *entity.Class) (err error) {
	err = c.db.WithContext(ctx).Delete(&entity).Error
	return
}
