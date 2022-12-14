package container

import (
	"technical-test/internal/domain/repository"
	"technical-test/internal/interfaces/infra/mysql"
	"technical-test/internal/interfaces/infra/seeder"
	"technical-test/internal/interfaces/usecase/class"
	"technical-test/internal/interfaces/usecase/student"
)

type Container struct {
	StudentService student.StudentService
	ClassService   class.ClassService
}

func (c *Container) Validate() *Container {
	if c.StudentService == nil {
		panic("studentService is nil")
	}
	if c.ClassService == nil {
		panic("classService is nil")
	}
	return c
}

func New() *Container {
	//setup db
	db := mysql.NewMysqlConnection()

	//setup Repository
	studentRepo := repository.NewStudentRepository(db)
	classRepo := repository.NewClassRepository(db)
	// redis := redis.NewRedisConnection().Setup()

	//setup Service
	studentService := student.NewService().SetStudentRepository(studentRepo).Validate()
	classService := class.NewService().SetClassRepository(classRepo).Validate()

	//create student seeder
	if err := seeder.CreateStudentSeeder(db); err != nil {
		panic("can't initialize student seeder")
	}

	return (&Container{
		StudentService: studentService,
		ClassService:   classService,
	}).Validate()
}
