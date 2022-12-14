package repository

import (
	"context"
	"technical-test/internal/domain/entity"

	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

type Student interface {
	Create(ctx context.Context, entity *entity.Student) (err error)
	GetAll(ctx context.Context, entity *[]entity.Student) (err error)
	GetById(ctx context.Context, id int64, entity *entity.Student) (err error)
	Update(ctx context.Context, entity *entity.Student) (err error)
	Delete(ctx context.Context, entity *entity.Student) (err error)
	GetStudentIdByName(ctx context.Context, name string) (id int64, err error)
}

func NewStudentRepository(db *gorm.DB) Student {
	if db == nil {
		panic("NewStudentRepository is nil")
	}
	return &studentRepository{db}
}

func (s *studentRepository) Create(ctx context.Context, entity *entity.Student) (err error) {
	err = s.db.WithContext(ctx).Create(&entity).Error
	return
}

func (s *studentRepository) GetAll(ctx context.Context, entity *[]entity.Student) (err error) {
	err = s.db.WithContext(ctx).Find(&entity).Error
	return
}

func (s *studentRepository) GetById(ctx context.Context, id int64, entity *entity.Student) (err error) {
	err = s.db.WithContext(ctx).Where(`id=?`, id).Find(&entity).Error
	return
}

func (s *studentRepository) Update(ctx context.Context, entity *entity.Student) (err error) {
	err = s.db.WithContext(ctx).Save(&entity).Error
	return
}

func (s *studentRepository) Delete(ctx context.Context, entity *entity.Student) (err error) {
	err = s.db.WithContext(ctx).Delete(&entity).Error
	return
}

func (s *studentRepository) GetStudentIdByName(ctx context.Context, name string) (id int64, err error) {
	model := entity.Student{}
	if err = s.db.WithContext(ctx).Where(`name=?`, name).Find(&model).Error; err != nil {
		return
	}
	return model.ID, nil
}
