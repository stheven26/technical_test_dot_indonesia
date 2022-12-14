package seeder

import (
	"technical-test/internal/domain/entity"
	"technical-test/internal/interfaces/infra/faker"

	"gorm.io/gorm"
)

type Student struct {
	Student *entity.Student
}

func CreateStudentSeeder(db *gorm.DB) error {
	std := faker.StudentFaker()
	student := []Student{
		{
			Student: std,
		},
	}
	for _, v := range student {
		err := db.Debug().Table("students").Create(&v.Student).Error
		if err != nil {
			return err
		}
	}
	return nil
}
