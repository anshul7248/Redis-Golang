package redis

import (
	"context"
	models "redis-apis/model"
)

type RedisUsecaseInterface interface {
	GetData(ctx context.Context) (*models.APIResponse, error)
}
