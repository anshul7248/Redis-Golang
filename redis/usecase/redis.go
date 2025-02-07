package usecase

import (
	"context"
	models "redis-apis/model"
	redisss "redis-apis/redis"
)

// import "redis-apis/redis/reposi tory"

type RedisUsecase struct {
	repository redisss.RedisRepoInterface
}

func NewRedisUsecase(repo redisss.RedisRepoInterface) redisss.RedisUsecaseInterface {
	return &RedisUsecase{
		repository: repo,
	}
}

func (r *RedisUsecase) GetData(ctx context.Context) (*models.APIResponse, error) {
	return r.repository.GetData(ctx)
}
