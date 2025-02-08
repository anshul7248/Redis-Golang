package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	db "redis-apis/adapter"
	"redis-apis/config"
	redisss "redis-apis/redis"
	"time"

	rediss "github.com/redis/go-redis/v9"

	models "redis-apis/model"

	"gorm.io/gorm"
)

type RedisRepository struct {
	DBConn    *gorm.DB
	RedisConn *rediss.Client
}

func DeleteKey(ctx context.Context, redisConn *rediss.Client, key string, value interface{}) (*models.APIResponse, error) {

	exists, err := redisConn.Exists(ctx, key).Result()
	if err != nil {
		fmt.Println("Redis key not found", err)
	}
	if exists == 1 {
		numDelete, err := redisConn.Del(ctx, key).Result()
		if err != nil {
			fmt.Println("Some error in deleting the key", err)
		}

		if numDelete > 0 {
			fmt.Println("Delete key from redis--->>>", key)
		} else {
			fmt.Println("Key not found in redis cache")
		}

		fmt.Println("Deleted key from redis cache--->>", key)
	}
	return nil, err
}

func CacheData(ctx context.Context, redisConn *rediss.Client, key string, value interface{}) (*models.APIResponse, error) {
	if value == nil {
		cacheData, err := redisConn.Get(ctx, key).Result()
		if err != nil {
			if err == rediss.Nil {
				log.Println("Cache Miss for key", key)
				return nil, nil
			}
			return nil, err
		}

		var cachedResponse models.APIResponse
		if err := json.Unmarshal([]byte(cacheData), &cachedResponse); err != nil {
			fmt.Println("Error parsing cached JSON:", err)
			return nil, err
		}
		log.Println("Returning cached response for key:  ", key)
		return &cachedResponse, nil
	}
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	err = redisConn.Set(ctx, key, string(jsonValue), 1*time.Minute).Err()
	if err != nil {
		return nil, err
	}
	log.Println("Cached Data for new key", key)
	return nil, nil
}

func NewRedisRepository(conn *gorm.DB, conf *config.Config) redisss.RedisRepoInterface {
	return &RedisRepository{
		DBConn:    conn,
		RedisConn: db.ConnectRedis(),
	}
}

func (r *RedisRepository) GetData(ctx context.Context) (*models.APIResponse, error) {
	mockDataModel := []models.MockData{}

	redisKey := "all_data"

	DeleteKey(ctx, r.RedisConn, redisKey, nil)
	cachedResponse, err := CacheData(ctx, r.RedisConn, redisKey, nil)
	if err != nil {
		fmt.Println("Error in getting cached data--->>>>", err)
	}
	if cachedResponse != nil {
		return cachedResponse, nil
	}
	if err := r.DBConn.Table("mock_data").Scan(&mockDataModel).Error; err != nil {
		return &models.APIResponse{Status: 0, Message: "Error", Result: 401}, nil
	}

	finalresponse := &models.APIResponse{Status: 0, Message: "Error", Result: 401, MockData: mockDataModel}
	if _, err := CacheData(ctx, r.RedisConn, redisKey, finalresponse); err != nil {
		fmt.Println("Error in caching data", err)
	} else {
		fmt.Println("Suceessfull", finalresponse)
	}
	return finalresponse, nil
}
