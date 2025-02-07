package main

import (
	"fmt"
	db "redis-apis/adapter"
	"redis-apis/config"
	"redis-apis/redis/controller"
	"redis-apis/redis/repository"

	// "redis-apis/redis/repository"
	"redis-apis/redis/usecase"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var onceRest sync.Once

func main() {
	onceRest.Do(func() {
		e := echo.New()

		config := config.GetConfig()

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		}))

		dbPsql := db.DB(config)

		redisClient := db.ConnectRedis()

		fmt.Println("Redis Connected on,", redisClient.Options().Addr)

		Repo := repository.NewRedisRepository(dbPsql, config)
		Uc := usecase.NewRedisUsecase(Repo)
		controller.NewRedisController(e, Uc)

		fmt.Println("Listening on port" + config.HttpConfig.HostPort)

		if err := e.Start(config.HttpConfig.HostPort); err != nil {
			// if err := e.StartTLS(config.HttpConfig.HostPort, config.HttpConfig.HostCert, config.HttpConfig.HostKey); err != nil {
			fmt.Println("Could not connect")
		}

	})
}
