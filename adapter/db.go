package db

import (
	"context"
	"fmt"
	"redis-apis/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ctx = context.Background()

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis connected ", pong)
	return client
}

func DB(config *config.Config) *gorm.DB {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Pass, config.Database.DBName)

	var err error

	DBPsql, err := gorm.Open(postgres.Open(connString))
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return DBPsql
}
