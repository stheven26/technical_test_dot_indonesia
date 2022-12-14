package student

import (
	"context"
	"strconv"
	"technical-test/internal/domain/entity"
	"technical-test/internal/domain/repository"
	"technical-test/internal/interfaces/infra/redis"
	"time"

	"technical-test/pkg/constants"

	"technical-test/pkg/log"
)

type service struct {
	studentRepository repository.Student
	redis             redis.RedisWrapper
}

func NewService() *service {
	return &service{}
}

func (s *service) SetStudentRepository(repo repository.Student) *service {
	s.studentRepository = repo
	return s
}
func (s *service) SetRedis(redis redis.RedisWrapper) *service {
	s.redis = redis
	return s
}

func (s *service) Validate() *service {
	if s.studentRepository == nil {
		panic("studentRepository is nil")
	}
	// if s.redis == nil {
	// 	panic("redis is nil")
	// }
	return s
}

func (s *service) CreateStudent(ctx context.Context, req StudentRequest) (res StudentResponse, err error) {
	data := entity.Student{
		Name:   req.Name,
		Age:    req.Age,
		School: req.School,
		Grade:  req.Grade,
	}
	if err = s.studentRepository.Create(ctx, &data); err != nil {
		log.Error(ctx, "createStudent Error", err.Error())
		return
	}
	res = StudentResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data: StudentResponseData{
			Id:        data.ID,
			Name:      data.Name,
			Age:       data.Age,
			School:    data.School,
			Grade:     data.Grade,
			CreatedAt: data.CreatedAt.Format(time.RFC3339),
			UpdatedAt: data.UpdatedAt.Format(time.RFC3339),
		},
	}
	return
}

func (s *service) GetAll(ctx context.Context) (res StudentResponse, err error) {
	student := []entity.Student{}
	if err = s.studentRepository.GetAll(ctx, &student); err != nil {
		log.Error(ctx, "getAllStudent error", err.Error())
		return
	}
	val := s.redis.GetAllStudent(ctx, constants.REDIS_STUDENT)
	if val == nil {
		err = s.redis.Set(ctx, constants.REDIS_STUDENT, 0, student)
		if err != nil {
			log.Error(ctx, "getRedisAllStudent error", err.Error())
			return
		}
	}
	res = StudentResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    val,
	}
	return
}

func (s *service) GetOneById(ctx context.Context, id string) (res StudentResponse, err error) {
	entity := entity.Student{}
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error(ctx, "convert to int64 Error", err.Error())
	}
	if err = s.studentRepository.GetById(ctx, newId, &entity); err != nil {
		log.Error(ctx, "getStudentById error", err.Error())
		return
	}
	res = StudentResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    entity,
	}
	return
}

func (s *service) UpdateById(ctx context.Context, id string, req StudentRequest) (res StudentResponse, err error) {
	entity := entity.Student{}
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error(ctx, "convert to int64 Error", err.Error())
	}
	if err = s.studentRepository.GetById(ctx, newId, &entity); err != nil {
		log.Error(ctx, "getStudentById error", err.Error())
		return
	}
	entity.Name = req.Name
	entity.Age = req.Age
	entity.School = req.School
	entity.Grade = req.Grade
	if err = s.studentRepository.Update(ctx, &entity); err != nil {
		log.Error(ctx, "updateStudent error", err.Error())
		return
	}
	res = StudentResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data: StudentResponseData{
			Id:        entity.ID,
			Name:      entity.Name,
			Age:       entity.Age,
			Grade:     entity.Grade,
			CreatedAt: entity.CreatedAt.Format(time.RFC3339),
			UpdatedAt: entity.UpdatedAt.Format(time.RFC3339),
		},
	}
	return
}

func (s *service) PatchById(ctx context.Context, id string, req PatchStudentNameRequest) (res StudentResponse, err error) {
	entity := entity.Student{}
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error(ctx, "convert to int64 Error", err.Error())
	}
	if err = s.studentRepository.GetById(ctx, newId, &entity); err != nil {
		log.Error(ctx, "getStudentById error", err.Error())
		return
	}
	entity.Name = req.Name
	if err = s.studentRepository.Update(ctx, &entity); err != nil {
		log.Error(ctx, "updateStudent error", err.Error())
		return
	}
	res = StudentResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data: StudentResponseData{
			Id:        entity.ID,
			Name:      entity.Name,
			CreatedAt: entity.CreatedAt.Format(time.RFC3339),
			UpdatedAt: entity.UpdatedAt.Format(time.RFC3339),
		},
	}
	return
}

func (s *service) DeleteById(ctx context.Context, id string) (res StudentResponse, err error) {
	entity := entity.Student{}
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error(ctx, "convert to int64 Error", err.Error())
	}
	if err = s.studentRepository.GetById(ctx, newId, &entity); err != nil {
		log.Error(ctx, "getStudentById error", err.Error())
		return
	}
	if err = s.studentRepository.Delete(ctx, &entity); err != nil {
		log.Error(ctx, "deleteStudentById error", err.Error())
		return
	}
	res = StudentResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
	}

	return
}
