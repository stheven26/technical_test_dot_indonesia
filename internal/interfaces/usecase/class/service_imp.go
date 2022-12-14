package class

import (
	"context"
	"strconv"
	"technical-test/internal/domain/entity"
	"technical-test/internal/domain/repository"
	"technical-test/internal/interfaces/infra/mysql"
	"technical-test/internal/interfaces/infra/redis"
	"technical-test/pkg/constants"
	"technical-test/pkg/log"
	"time"
)

type service struct {
	classRepository repository.Class
	redis           redis.RedisWrapper
}

func NewService() *service {
	return &service{}
}

func (s *service) SetClassRepository(repo repository.Class) *service {
	s.classRepository = repo
	return s
}
func (s *service) SetRedis(redis redis.RedisWrapper) *service {
	s.redis = redis
	return s
}

func (s *service) Validate() *service {
	if s.classRepository == nil {
		panic("classRepository is nil")
	}
	if s.redis == nil {
		panic("redis is nil")
	}
	return s
}

func (s *service) CreateClass(ctx context.Context, req ClassRequest) (res ClassResponse, err error) {
	db := mysql.NewMysqlConnection()
	std := repository.NewStudentRepository(db)
	id, err := std.GetStudentIdByName(ctx, req.StudentName)
	if err != nil {
		log.Error(ctx, "getStudentIdByName Error", err.Error())
		return
	}
	data := entity.Class{
		StudentID:     id,
		Subject:       req.Subject,
		DurationClass: req.DurationClass,
	}
	if err = s.classRepository.Create(ctx, &data); err != nil {
		log.Error(ctx, "createClass Error", err.Error())
		return
	}
	res = ClassResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data: ClassResponseData{
			Id:            data.ID,
			Subject:       data.Subject,
			DurationClass: data.DurationClass,
			CreatedAt:     data.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     data.UpdatedAt.Format(time.RFC3339),
		}}
	return
}

func (s *service) GetAll(ctx context.Context) (res ClassResponse, err error) {
	class := []entity.Class{}
	if err = s.classRepository.GetAll(ctx, &class); err != nil {
		log.Error(ctx, "getAllClass error", err.Error())
		return
	}
	val := s.redis.GetAllClass(ctx, constants.REDIS_CLASS)
	if val == nil {
		err = s.redis.Set(ctx, constants.REDIS_CLASS, 0, class)
		if err != nil {
			log.Error(ctx, "getRedisAllClass error", err.Error())
			return
		}
	}
	res = ClassResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    val,
	}
	return
}

func (s *service) GetOneById(ctx context.Context, id string) (res ClassResponse, err error) {
	entity := entity.Class{}
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error(ctx, "convert to int64 Error", err.Error())
	}
	if err = s.classRepository.GetById(ctx, newId, &entity); err != nil {
		log.Error(ctx, "getClassById error", err.Error())
		return
	}
	res = ClassResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    entity,
	}
	return
}
func (s *service) DeleteById(ctx context.Context, id string) (res ClassResponse, err error) {
	entity := entity.Class{}
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error(ctx, "convert to int64 Error", err.Error())
	}
	if err = s.classRepository.GetById(ctx, newId, &entity); err != nil {
		log.Error(ctx, "getClassById error", err.Error())
		return
	}
	if err = s.classRepository.Delete(ctx, &entity); err != nil {
		log.Error(ctx, "deleteClassById error", err.Error())
		return
	}
	res = ClassResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
	}
	return
}
