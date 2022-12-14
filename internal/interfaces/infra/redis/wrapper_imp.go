package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"technical-test/internal/domain/entity"
	"technical-test/pkg/config"
	"time"

	"github.com/go-redis/redis/v9"
)

type redisWrapper struct {
	redis *redis.Client
}

func NewRedisConnection() *redisWrapper {
	return &redisWrapper{}
}

func (r *redisWrapper) Setup() RedisWrapper {
	loadEnv := config.LoadEnv()
	redisEnv := map[string]string{
		"host":     loadEnv.GetString("REDIS_HOST"),
		"username": loadEnv.GetString("REDIS_USERNAME"),
		"password": loadEnv.GetString("REDIS_PASSWORD"),
		"db":       loadEnv.GetString("REDIS_DB"),
	}
	db, err := strconv.Atoi(redisEnv["db"])
	if err != nil {
		fmt.Println(err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisEnv["host"],
		Username: redisEnv["username"],
		Password: redisEnv["password"],
		DB:       db,
	})
	r.redis = redisClient
	return r
}

func (r *redisWrapper) GetAllStudent(ctx context.Context, key string) (resp []entity.Student) {
	client := r.redis
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return
	}
	if err = json.Unmarshal([]byte(val), &resp); err != nil {
		return
	}
	return
}

func (r *redisWrapper) GetAllClass(ctx context.Context, key string) (resp []entity.Class) {
	client := r.redis
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return
	}
	if err = json.Unmarshal([]byte(val), &resp); err != nil {
		return
	}
	return
}

func (r *redisWrapper) Set(ctx context.Context, key string, expirationTime time.Duration, req interface{}) (err error) {
	client := r.redis
	json, err := json.Marshal(req)
	if err != nil {
		err = fmt.Errorf("redis marshal error: %w", err)
		return
	}
	err = client.Set(ctx, key, json, expirationTime).Err()
	if err != nil {
		err = fmt.Errorf("redis set error: %w", err)
		return
	}
	return
}
