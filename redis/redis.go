package redis

import (
	"context"
	models "redis-apis/model"
)

type RedisRepoInterface interface {
	GetData(ctx context.Context) (*models.APIResponse, error)
}
