package redis

import (
	"context"
	"technical-test/internal/domain/entity"
	"time"
)

type RedisWrapper interface {
	GetAllStudent(ctx context.Context, key string) (resp []entity.Student)
	GetAllClass(ctx context.Context, key string) (resp []entity.Class)
	Set(ctx context.Context, key string, expirationTime time.Duration, req interface{}) (err error)
}
